package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "player1/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

type Player struct {
    ID                 string
    GameModePreference string
}

func player_status(client pb.MatchMakerClient, playerID string) {

    // Ciclo para consultar el estado del jugador y permitirle salir de la cola si lo desea
    for {

        // Mostrar opciones al jugador
        var estado int
        fmt.Println("\n1. Consultar estado / 0. Salir de la Queue")
        fmt.Print("Ingrese su opción: ")
        _, err := fmt.Scanln(&estado)
        if err != nil {
            fmt.Println("Error leyendo la opción, intente de nuevo.")
            return
        }

        if estado == 1 {
            // Consultar Estado del Jugador
            ctxStatus, cancelStatus := context.WithTimeout(context.Background(), time.Second)
            defer cancelStatus()

            statusResp, err := client.GetPlayerStatus(ctxStatus, &pb.PlayerStatusRequest{PlayerId: playerID})
            if err != nil {
                log.Fatalf("Error en GetPlayerStatus: %v", err)
            }
            fmt.Printf("Estado del jugador: %s\n", statusResp.Status.String())

        } else if estado == 0 {
            // Salir de la Queue
            ctxLeave, cancelLeave := context.WithTimeout(context.Background(), time.Second)
            defer cancelLeave()

            _, err := client.LeaveQueuePlayer(ctxLeave, &pb.PlayerInfoRequest{PlayerId: playerID})
            if err != nil {
                log.Fatalf("Error en LeaveQueue: %v", err)
            }
            fmt.Printf("Se ha retirado al jugador de la cola\n")
            break

        } else {
            // Opciones inválidas
            fmt.Println("Opción no válida, por favor intente de nuevo.")

        }
    }
}

func player_queue(client pb.MatchMakerClient, playerID string) {

    // Enviar solicitud para poner al jugador en la cola
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Poner al jugador en la cola
    queueResp, err := client.QueuePlayer(ctx, &pb.PlayerInfoRequest{PlayerId: playerID})
    if err != nil {
        log.Fatalf("Error en QueuePlayer: %v", err)
    }
    fmt.Printf("%s\n", queueResp.Message)

    // Consultar el estado del jugador
    player_status(client, playerID)

}

func main() {
    // Conectar al servidor MatchMaker
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar: %v", err)
    }
    defer conn.Close()

    // Crear un cliente MatchMaker
    client := pb.NewMatchMakerClient(conn)

    // Datos de Player 1
    player := Player{
        ID: "001",
        GameModePreference: "clasico",
    }

    options := 0
    for {

		fmt.Println("\n¿Quiere enlistar el jugador en la Queue? (1. Sí / 0. No)")

		// Leer la opción del usuario
		fmt.Print("Ingrese su opción: ")
		_, err := fmt.Scanln(&options)
		if err != nil {
			fmt.Println("Error leyendo la opción, intente de nuevo.")
			continue
		}

		// Ejecutar la opción seleccionada
		if options == 1 {
			// Poner a Jugador en Queue
			player_queue(client, player.ID)

		} else if options == 0 {
            // Terminar el programa
			break

		}else {
			// Opción no válida
			fmt.Println("\nOpción no válida, por favor intente de nuevo.")

		}
	}
}
