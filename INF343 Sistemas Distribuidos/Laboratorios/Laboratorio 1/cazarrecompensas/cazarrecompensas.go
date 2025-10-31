package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cazarrecompensas/proto/grpc-server/proto"
)

func main() {
	conn, err := grpc.Dial("gobierno:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Gobierno Mundial: %v", err)
	}
	defer conn.Close()

	clienteGobierno := pb.NewGobiernoMundialClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var nombre string
	fmt.Print("Introduce el nombre del cazador: ")
	fmt.Scanln(&nombre)

	res, err := clienteGobierno.RegistrarCazarrecompensas(ctx, &pb.EstadoCazarrecompensas{
		NombreCazarrecompensas: nombre,
		Reputacion:             100.0,
	})
	if err != nil {
		log.Fatalf("Error al registrar cazarrecompensas: %v", err)
	}
	if !res.Exito {
		log.Fatalf("Registro fallido: %s", res.Mensaje)
	}
	log.Printf("Registro exitoso: %s\n", res.Mensaje)

	// Inicializar contadores de entregas consecutivas
	var ventasConsecutivasSubmundo, ventasConsecutivasMarina int

	for {
		// Obtener la lista actualizada de piratas buscados
		lista, err := obtenerListaPiratas(clienteGobierno)
		if err != nil {
			log.Fatalf("Error al consultar piratas: %v", err)
		}

		if len(lista.Piratas) == 0 {
			log.Println("No hay piratas buscados actualmente.")
			break
		}

		// Buscar el primer pirata con estado BUSCADO
		var pirata *pb.Pirata
		for _, p := range lista.Piratas {
			if p.Estado == pb.EstadoPirata_BUSCADO {
				pirata = p
				break
			}
		}

		if pirata == nil {
			log.Println("No se encontraron piratas en estado BUSCADO.")
			time.Sleep(3 * time.Second)
			continue
		}

		// Imprimir datos del pirata seleccionado
		log.Printf("ID: %s, Nombre: %s, Recompensa: %d, Peligrosidad: %v, Estado: %v\n",
			pirata.Id, pirata.Nombre, pirata.Recompensa, pirata.Peligrosidad, pirata.Estado)

		// Evaluar si el pirata se escapa (antes de decidir el destino)
		if evaluarFuga(pirata.Peligrosidad, pirata.Nombre) {
			log.Println("El pirata se ha escapado.")
			log.Println("────────────────────────────")
			time.Sleep(2 * time.Second)
			continue
		}

		_, err = clienteGobierno.ActualizarEstadoPirata(context.Background(), &pb.Pirata{
			Id:     pirata.Id,
			Estado: pb.EstadoPirata_EN_CAMINO,
		})
		if err != nil {
			log.Printf("Error al notificar EN_CAMINO: %v", err)
			continue
		}
		log.Printf("Gobierno notificado: %s ahora está EN_CAMINO", pirata.Nombre)

		// Decidir destino para el pirata (Marina o Submundo)
		destino := decidirDestino(pirata, ventasConsecutivasSubmundo)

		entrega := &pb.EntregaPirata{
			NombreCazarrecompensas: nombre,
			Pirata:                 pirata,
			Reputacion:             100.0, // esto podrías consultar en tiempo real
			Metodo:                 destino,
		}

		// Enviar al destino
		var resultado *pb.ResultadoEntrega
		if destino == pb.MetodoEntrega_MARINA {
			resultado, err = enviarAMarina(entrega)
		} else {
			resultado, err = enviarAlSubmundo(entrega)
		}

		if err != nil {
			log.Printf("Error al enviar pirata a %v: %v", destino, err)
			time.Sleep(2 * time.Second)
			continue
		}

		procesarResultadoEntrega(clienteGobierno, resultado, pirata)

		actualizarContadores(&ventasConsecutivasSubmundo, &ventasConsecutivasMarina, destino)

		log.Println("────────────────────────────")

		time.Sleep(2 * time.Second)
	}

}

// Función para obtener la lista de piratas desde el Gobierno Mundial
func obtenerListaPiratas(clienteGobierno pb.GobiernoMundialClient) (*pb.ListaPiratas, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Consultar lista de piratas buscados
	return clienteGobierno.ListarPiratasBuscados(ctx, &emptypb.Empty{})
}

// Función para decidir el destino del pirata (Submundo o Marina)
func decidirDestino(pirata *pb.Pirata, ventasConsecutivasSubmundo int) pb.MetodoEntrega {
    var destino pb.MetodoEntrega

    // Si la reputación del cazarrecompensas es baja (<= 50), se fuerza al Submundo
    if pirata.Recompensa <= 50 {
        destino = pb.MetodoEntrega_SUBMUNDO
        log.Printf("Reputación baja. Se fuerza entrega al Submundo para %s.", pirata.Nombre)
    } else {
        // Calcular riesgos
        riesgoMarina := 0.25
        if pirata.Peligrosidad == pb.Peligrosidad_ALTA {
            riesgoMarina = 0.35
        }

        riesgoSubmundo := 0.35
        if ventasConsecutivasSubmundo >= 3 {
            riesgoSubmundo = 1 - (1-0.35)*(1-0.25)
        }

        // Lógica de decisión entre Submundo y Marina
        if riesgoMarina < riesgoSubmundo {
            destino = pb.MetodoEntrega_MARINA
        } else {
            destino = pb.MetodoEntrega_SUBMUNDO
        }
    }

    return destino
}

// Función para enviar el pirata a la Marina
func enviarAMarina(entrega *pb.EntregaPirata) (*pb.ResultadoEntrega, error) {
	connMarina, err := grpc.Dial("marina:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con la Marina: %v", err)
		return nil, err
	}
	defer connMarina.Close()

	clienteMarina := pb.NewMarinaClient(connMarina)

	// Enviar el pirata a la Marina
	resultado, err := clienteMarina.RecibirPirata(context.Background(), entrega)
	if err != nil {
		log.Printf("Error al enviar pirata a la Marina: %v", err)
		return nil, err
	}

	return resultado, nil
}

// Función para enviar el pirata al Submundo
func enviarAlSubmundo(entrega *pb.EntregaPirata) (*pb.ResultadoEntrega, error) {
	connSubmundo, err := grpc.Dial("submundo:50053", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error al conectar con el Submundo: %v", err)
		return nil, err
	}
	defer connSubmundo.Close()

	clienteSubmundo := pb.NewSubmundoClient(connSubmundo)

	// Enviar el pirata al Submundo
	resultado, err := clienteSubmundo.ComprarPirata(context.Background(), entrega)
	if err != nil {
		log.Printf("Error al enviar pirata al Submundo: %v", err)
		return nil, err
	}
	return resultado, nil
}

// Función para procesar el resultado de la entrega
func procesarResultadoEntrega(clienteGobierno pb.GobiernoMundialClient, resultado *pb.ResultadoEntrega, pirata *pb.Pirata) {
	if resultado != nil && resultado.Aceptado {
		_, err := clienteGobierno.ActualizarEstadoPirata(context.Background(), &pb.Pirata{
			Id:     pirata.Id,
			Estado: pb.EstadoPirata_ENTREGADO,
		})
		if err != nil {
			log.Printf("Error al actualizar estado ENTREGADO para %s: %v", pirata.Nombre, err)
		}
		log.Printf("El pirata %s fue entregado correctamente.", pirata.Nombre)
	} else {
		log.Printf("La entrega del pirata %s fue rechazada.", pirata.Nombre)
	}
}

// Función para actualizar los contadores de entregas consecutivas
func actualizarContadores(ventasConsecutivasSubmundo, ventasConsecutivasMarina *int, destino pb.MetodoEntrega) {
	switch destino {
	case pb.MetodoEntrega_SUBMUNDO:
		(*ventasConsecutivasSubmundo)++
		*ventasConsecutivasMarina = 0
		log.Printf("Entrega al Submundo (consecutivas: %d)", *ventasConsecutivasSubmundo)
	case pb.MetodoEntrega_MARINA:
		(*ventasConsecutivasMarina)++
		*ventasConsecutivasSubmundo = 0
		log.Printf("Entrega a la Marina (consecutivas: %d)", *ventasConsecutivasMarina)
	}
}

// Función para simular una probabilidad aleatoria
func randomChance(prob float64) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() < prob
}

// Función para evaluar si un pirata se escapa durante el transporte
func evaluarFuga(peligrosidad pb.Peligrosidad, nombre string) bool {
	var prob float64
	switch peligrosidad {
	case pb.Peligrosidad_BAJA:
		prob = 0.15
	case pb.Peligrosidad_MEDIA:
		prob = 0.25
	case pb.Peligrosidad_ALTA:
		prob = 0.45
	}
	if randomChance(prob) {
		log.Printf("El pirata %s se fugó durante el transporte.", nombre)
		return true
	}
	return false
}
