package main

import (
    "context"
    "log"
    "net"

    pb "matchmaker/proto/grpc-server/proto"
    "google.golang.org/grpc"
)

/*
Structs:
- Player: Representa un jugador en el sistema de matchmaking.
*/

type Player struct {
    ID                 string
    status             string
}

/*
Listas, Variables Globales:
- colas: Lista de jugadores en cola (1v1 por defecto).
*/

var colas = make([]*Player, 2)

type matchMakerServer struct {
    pb.UnimplementedMatchMakerServer
}

/*
Funciones gRPC:
- QueuePlayer: Agrega un jugador a la cola de matchmaking.
- LeaveQueuePlayer: Permite a un jugador salir de la cola.
- GetPlayerStatus: Consulta el estado de un jugador en la cola.
*/
func (s *matchMakerServer) QueuePlayer(ctx context.Context, req *pb.PlayerInfoRequest) (*pb.QueuePlayerResponse, error) {
	
	player := &Player{
		ID:     req.PlayerId,
		status: "BUSCANDO PARTIDA",
	}
	colas = append(colas, player)

    return &pb.QueuePlayerResponse{
        StatusCode: pb.StatusCode_EXITO,
        Message:    "Jugador en Cola",
    }, nil
}

func (s *matchMakerServer) LeaveQueuePlayer(ctx context.Context, req *pb.PlayerInfoRequest) (*pb.Empty, error) {
    for i, player := range colas {
        if player != nil && player.ID == req.PlayerId {
            colas = append(colas[:i], colas[i+1:]...)
            return &pb.Empty{}, nil
        }
    }
	// Si no se encuentra el jugador
    return &pb.Empty{}, nil
}

func (s *matchMakerServer) GetPlayerStatus(ctx context.Context, req *pb.PlayerStatusRequest) (*pb.PlayerStatusResponse, error) {
    for _, player := range colas {
        if player != nil && player.ID == req.PlayerId {
            var status pb.PlayerStatus
            switch player.status {
            case "LIBRE":
                status = pb.PlayerStatus_LIBRE
            case "BUSCANDO PARTIDA":
                status = pb.PlayerStatus_BUSCANDO_PARTIDA
            case "EN PARTIDA":
                status = pb.PlayerStatus_EN_PARTIDA
            }
            return &pb.PlayerStatusResponse{
                Status:      status,
                MatchId:     "001", // HARDCODEADO
                ServerClock: "00:00:00", // HARDCODEADO
            }, nil
        }
    }
    // Si no se encuentra el jugador
    return &pb.PlayerStatusResponse{
        Status:      pb.PlayerStatus_LIBRE,
        MatchId:     "001", // HARDCODEADO
		ServerClock: "00:00:00", // HARDCODEADO
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
