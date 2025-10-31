package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "player1/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

type Player struct {
	ID                 string
	GameModePreference string
}

func player_status(client pb.MatchMakerClient, playerID string, scanner *bufio.Scanner) {

	for {
		// Mostrar opciones al jugador
		fmt.Println("\n1. Consultar estado / 0. Salir de la Queue")
		fmt.Print("Ingrese su opción: ")

		if !scanner.Scan() {
			fmt.Println("Error leyendo la opción, intente de nuevo.")
			return
		}

		input := strings.TrimSpace(scanner.Text())
		estado, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Por favor ingrese un número válido.")
			continue
		}

		if estado == 1 {
			ctxStatus, cancelStatus := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelStatus()

			statusResp, err := client.GetPlayerStatus(ctxStatus, &pb.PlayerStatusRequest{PlayerId: playerID})
			if err != nil {
				log.Printf("Error en GetPlayerStatus: %v", err)
				fmt.Println("No se pudo consultar el estado. Intente de nuevo.")
				continue
			}

			var statusText string
			switch statusResp.Status {
			case pb.PlayerStatus_LIBRE:
				statusText = "LIBRE"
			case pb.PlayerStatus_BUSCANDO_PARTIDA:
				statusText = "BUSCANDO PARTIDA"
			case pb.PlayerStatus_EN_PARTIDA:
				statusText = "EN PARTIDA"
			default:
				statusText = "DESCONOCIDO"
			}

			fmt.Printf("Estado del jugador: %s\n", statusText)
			if statusResp.MatchId != "" && statusResp.MatchId != "001" {
				fmt.Printf("Match ID: %s\n", statusResp.MatchId)
			}

		} else if estado == 0 {
			// Salir de la Queue
			ctxLeave, cancelLeave := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelLeave()

			_, err := client.LeaveQueuePlayer(ctxLeave, &pb.PlayerInfoRequest{PlayerId: playerID})
			if err != nil {
				log.Printf("Error en LeaveQueue: %v", err)
				fmt.Println("Error al intentar salir de la cola.")
			} else {
				fmt.Printf("Se ha retirado al jugador de la cola\n")
			}
			break

		} else {
			// Opciones inválidas
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}

func player_queue(client pb.MatchMakerClient, playerID string, gameMode string, scanner *bufio.Scanner) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	queueResp, err := client.QueuePlayer(ctx, &pb.PlayerInfoRequest{
		PlayerId:           playerID,
		GameModePreference: gameMode,
	})
	if err != nil {
		log.Printf("Error en QueuePlayer: %v", err)
		fmt.Println("No se pudo unir a la cola. Intente de nuevo.")
		return
	}

	if queueResp.StatusCode == pb.StatusCode_FALLO {
		fmt.Printf("Error: %s\n", queueResp.Message)
		return
	}

	fmt.Printf("%s\n", queueResp.Message)

	fmt.Println("Verificando estado inmediatamente...")
	ctxVerify, cancelVerify := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelVerify()

	statusResp, err := client.GetPlayerStatus(ctxVerify, &pb.PlayerStatusRequest{PlayerId: playerID})
	if err != nil {
		log.Printf("Error verificando estado: %v", err)
	} else {
		var statusText string
		switch statusResp.Status {
		case pb.PlayerStatus_BUSCANDO_PARTIDA:
			statusText = "BUSCANDO PARTIDA"
		case pb.PlayerStatus_LIBRE:
			statusText = "LIBRE"
		case pb.PlayerStatus_EN_PARTIDA:
			statusText = "EN PARTIDA"
		default:
			statusText = "DESCONOCIDO"
		}
		fmt.Printf("Estado confirmado: %s\n", statusText)
	}

	player_status(client, playerID, scanner)
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewMatchMakerClient(conn)

	player := Player{
		ID:                 "Player1",
		GameModePreference: "clasico",
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("=== Jugador %s iniciado ===\n", player.ID)
	fmt.Printf("Estado inicial: LIBRE\n")

	for {
		fmt.Println("\n=== MENÚ ===")
		fmt.Println("1. Unirse a cola de matchmaking")
		fmt.Println("2. Consultar estado actual")
		fmt.Println("0. Salir")
		fmt.Print("Ingrese su opción: ")

		if !scanner.Scan() {
			fmt.Println("Error leyendo entrada. Terminando programa.")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Por favor ingrese un número válido.")
			continue
		}

		switch option {
		case 1:
			// Poner a Jugador en Queue
			player_queue(client, player.ID, player.GameModePreference, scanner)

		case 2:
			// Consultar estado del Jugador
			ctxStatus, cancelStatus := context.WithTimeout(context.Background(), 5*time.Second)
			statusResp, err := client.GetPlayerStatus(ctxStatus, &pb.PlayerStatusRequest{PlayerId: player.ID})
			cancelStatus()

			if err != nil {
				log.Printf("Error consultando estado: %v", err)
				fmt.Println("No se pudo consultar el estado.")
			} else {
				var statusText string
				switch statusResp.Status {
				case pb.PlayerStatus_LIBRE:
					statusText = "LIBRE"
				case pb.PlayerStatus_BUSCANDO_PARTIDA:
					statusText = "BUSCANDO PARTIDA"
				case pb.PlayerStatus_EN_PARTIDA:
					statusText = "EN PARTIDA"
				default:
					statusText = "DESCONOCIDO"
				}
				fmt.Printf("Estado actual: %s\n", statusText)
			}

		case 0:
			// Terminar el programa
			fmt.Println("Cerrando aplicación...")
			return

		default:
			// Opción no válida
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}
