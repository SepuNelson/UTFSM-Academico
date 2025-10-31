package main

import (
	"context"
	"fmt"
	"log"

	pb "entrenador/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

/*
Variables and List Definitions:

*/

var registroParticipaciones []*pb.RegistroParticipante

/*
Local Functions Definitions:
- option_1: Función para mostrar los torneos disponibles.
- option_2: Función para inscribirse en un torneo.
- option_3: Función para ver las notificaciones recibidas.
- option_4: Función para ver el estado actual.
*/

// Función para mostrar los torneos disponibles
func option_1(client pb.LCPServiceClient) {

	fmt.Println("\n================================")
	fmt.Println("Torneos disponibles")

	// Crear Request
	req := &pb.Empty{}

	// Llamar al método remoto
	resp, err := client.PublicarTorneos(context.Background(), req)
	if err != nil {
		fmt.Printf("Error al obtener torneos: %v\n", err)
	} else {
		for _, t := range resp.Torneos {
			fmt.Printf("- ID: %d | Región: %s\n", t.Id, t.Region)
		}
	}

	fmt.Println("================================")
}

// Función para inscribirse en un torneo
func option_2(client pb.LCPServiceClient, entrenador *pb.Trainer) {

	fmt.Println("\n================================")

	// Variable para almacenar la ID del torneo
	fmt.Print("\nIngrese la ID torneo al que se quiere inscribir: ")
	var id_torneo int
	_, err := fmt.Scanln(&id_torneo)
	if err != nil {
		fmt.Println("Error leyendo la opción, intente de nuevo.")
		return
	}

	// Crear el request solo con el entrenador y el id del torneo
	req := &pb.InscripcionRequest{
		Entrenador: entrenador,
		IdTorneo:   int32(id_torneo),
	}

	// Llamar al método remoto
	resp, err := client.InscribirseTorneo(context.Background(), req)
	if err != nil {
		fmt.Printf("Error al inscribirse en el torneo: %v\n", err)

	} else if resp.Mensaje == "Inscripción exitosa.\n" {

		entrenador.Estado=pb.EstadoEntrenador_INSCRITO

		request := &pb.InscripcionRequest{
			Entrenador: entrenador,
			IdTorneo:   int32(id_torneo),
		}

		respuesta, err := client.RegistrarParticipacion(context.Background(), request)
		if err != nil {
			fmt.Printf("Error al registrar entrenador %v\n", err)
		}

		registroParticipaciones = append(registroParticipaciones, respuesta)
		fmt.Println(resp.Mensaje)
	} else {
		fmt.Println(resp.Mensaje)
	}

	fmt.Println("\n================================")
}

// Función para ver las notificaciones recibidas
func option_3( /*client pb.LCPServiceClient*/ ) {

	fmt.Println("\n================================")
	fmt.Println("Notificaciones recibidas")

	fmt.Println("================================")
}

// Función para ver el estado actual
func option_4(entrenador *pb.Trainer) {
	fmt.Println("\n================================")
	fmt.Println("Estado actual del Entrenador")
    fmt.Printf("Ranking: %d\n", entrenador.Estado)
	fmt.Println("================================")
}

// Función principal
func main() {

	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to LCP server: %v", err)
	}
	fmt.Println("Conexión exitosa con el servidor LCP.")
	defer conn.Close()

	// Crear cliente gRPC
	client := pb.NewLCPServiceClient(conn)

	// Crear perfil de Entrenador
	entrenador := &pb.Trainer{
		Id:         000,	
		Nombre:     "Nelson",
		Region:     "Kanto",
		Ranking:    1500,
		Estado:     pb.EstadoEntrenador_ACTIVO,
		Suspension: 0,
	}


	options := 0
	for {

		// Mostrar el menú
		fmt.Println("\n=========== M E N Ú ===========")
		fmt.Println("1. Consultar Torneos Disponibles")
		fmt.Println("2. Inscribirse en un Torneo")
		fmt.Println("3. Ver notificaciones recibidas")
		fmt.Println("4. Ver estado actual")
		fmt.Println("5. Salir")

		// Leer la opción del usuario
		fmt.Print("\nIngrese su opción: ")
		_, err := fmt.Scanln(&options)
		if err != nil {
			fmt.Println("Error leyendo la opción, intente de nuevo.")
			options = 0
			continue
		}

		fmt.Println("===============================")

		// Ejecutar la opción seleccionada
		if options == 1 {

			// Consultar torneos disponibles
			option_1(client)

		} else if options == 2 {

			// Inscribirse en un torneo
			option_2(client, entrenador)

		} else if options == 3 {

			// Ver notificaciones recibidas
			option_3( /*client*/ )

		} else if options == 4 {

			// Ver estado actual
			option_4(entrenador)

		} else if options == 5 {

			// Salir del programa
			fmt.Println("\nSaliendo del programa...")
			fmt.Printf("%v\n", registroParticipaciones)
			break

		} else {

			// Opción no válida
			fmt.Println("\nOpción no válida, por favor intente de nuevo.")

		}
	}

}
