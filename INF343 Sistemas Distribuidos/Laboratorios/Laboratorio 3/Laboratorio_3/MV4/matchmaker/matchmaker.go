package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "matchmaker/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

/*
Structs:
- Player: Representa un jugador en el sistema de matchmaking.
*/

type Player struct {
	ID                 string
	Status             string
	GameModePreference string
}

/*
Listas, Variables Globales:
- queue: Lista de jugadores en cola (sincronizada con mutex).
- playerStates: Mapa para mantener estado de todos los jugadores registrados.
*/

var (
	queue        []*Player
	playerStates = make(map[string]*Player)
	mutex        sync.Mutex
)

type matchMakerServer struct {
	pb.UnimplementedMatchMakerServer
}

/*
Funciones Locales:
- tryMatchPlayers: Intenta emparejar jugadores de la cola.
*/

func tryMatchPlayers() {

    for len(queue) >= 2 {
        // Tomar los dos primeros jugadores de la cola
        p1 := queue[0]
        p2 := queue[1]

        // Eliminar los dos jugadores de la cola
        queue = queue[2:]

        // Cambiar su estado a EN_PARTIDA
        if playerStates[p1.ID] != nil {
            playerStates[p1.ID].Status = "EN_PARTIDA"
        }
        if playerStates[p2.ID] != nil {
            playerStates[p2.ID].Status = "EN_PARTIDA"
        }

    }
}

/*
Funciones gRPC:
- QueuePlayer: Agrega un jugador a la cola de matchmaking.
- LeaveQueuePlayer: Permite a un jugador salir de la cola.
- GetPlayerStatus: Consulta el estado de un jugador en la cola.
*/
func (s *matchMakerServer) QueuePlayer(ctx context.Context, req *pb.PlayerInfoRequest) (*pb.QueuePlayerResponse, error) {

	mutex.Lock()
	defer mutex.Unlock()

	for _, p := range queue {
		if p.ID == req.PlayerId {
			return &pb.QueuePlayerResponse{
				StatusCode: pb.StatusCode_FALLO,
				Message:    "Jugador ya en cola",
			}, nil
		}
	}

	player := &Player{
		ID:                 req.PlayerId,
		Status:             "BUSCANDO_PARTIDA",
		GameModePreference: req.GameModePreference,
	}

	queue = append(queue, player)
	playerStates[req.PlayerId] = player

	log.Printf("[MatchMaker] Jugador %s en cola para modo '%s'", player.ID, player.GameModePreference)

	tryMatchPlayers()

	return &pb.QueuePlayerResponse{
		StatusCode: pb.StatusCode_EXITO,
		Message:    "Jugador en Cola",
	}, nil
}

func (s *matchMakerServer) LeaveQueuePlayer(ctx context.Context, req *pb.PlayerInfoRequest) (*pb.Empty, error) {

	mutex.Lock()
	defer mutex.Unlock()

	for i, player := range queue {
		if player != nil && player.ID == req.PlayerId {
			queue = append(queue[:i], queue[i+1:]...)

			if p, exists := playerStates[req.PlayerId]; exists {
				p.Status = "LIBRE"
			}

			log.Printf("[MatchMaker] Jugador %s ha salido de la cola", req.PlayerId)
			return &pb.Empty{}, nil
		}
	}

	log.Printf("[MatchMaker] Jugador %s no estaba en cola", req.PlayerId)
	return &pb.Empty{}, nil
}

func (s *matchMakerServer) GetPlayerStatus(ctx context.Context, req *pb.PlayerStatusRequest) (*pb.PlayerStatusResponse, error) {

	mutex.Lock()
	defer mutex.Unlock()

	if player, exists := playerStates[req.PlayerId]; exists {
		var status pb.PlayerStatus
		switch player.Status {
		case "LIBRE":
			status = pb.PlayerStatus_LIBRE
		case "BUSCANDO_PARTIDA":
			status = pb.PlayerStatus_BUSCANDO_PARTIDA
		case "EN_PARTIDA":
			status = pb.PlayerStatus_EN_PARTIDA
		default:
			status = pb.PlayerStatus_LIBRE
		}

		return &pb.PlayerStatusResponse{
			Status:      status,
			MatchId:     "",
			ServerClock: "00:00:00",
		}, nil
	}

	playerStates[req.PlayerId] = &Player{
		ID:     req.PlayerId,
		Status: "LIBRE",
	}

	return &pb.PlayerStatusResponse{
		Status:      pb.PlayerStatus_LIBRE,
		MatchId:     "",
		ServerClock: "00:00:00",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMatchMakerServer(grpcServer, &matchMakerServer{})

	log.Println("MatchMaker se está ejecutando correctamente...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
