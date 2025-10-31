package main

import (
	"context"
	"log"
	"net"
	"math/rand"

	"google.golang.org/grpc"
	pb "submundo/proto/grpc-server/proto"
)

// Definición del servidor Submundo
type servidorSubmundo struct {
	pb.UnimplementedSubmundoServer
}

func (s *servidorSubmundo) ComprarPirata(ctx context.Context, pirata *pb.EntregaPirata) (*pb.ResultadoEntrega, error) {
	// Se calcula un valor aleatorio entre 100% y 150% de la recompensa oficial
	precioCompra := float64(pirata.Pirata.Recompensa) * (1 + rand.Float64()*0.5)

	// Simular un fraude con un 35% de probabilidad
	if rand.Float64() < 0.35 {
		// Si hay fraude, el cazarrecompensas no recibe nada
		log.Printf("FRAUDE DETECTADO: La recompensa del pirata %s no fue entregada", pirata.Pirata.Nombre)
		return &pb.ResultadoEntrega{
			Aceptado:          false,
			RecompensaPagada: uint64(0),
		}, nil
	}

	// Si el pirata tiene una recompensa válida, se procede con la compra (sin fraude)
	return &pb.ResultadoEntrega{
		Aceptado:          true,
		RecompensaPagada: uint64(precioCompra),
	}, nil
}

func (s *servidorSubmundo) InterceptarTransporte(ctx context.Context, pirata *pb.Pirata) (*pb.ResultadoTransporte, error) {
	// Si el pirata tiene una recompensa mayor a 200 millones de Berries, mercenarios
	resultado := &pb.ResultadoTransporte{
		EstadoFinal:    pb.EstadoPirata_ENTREGADO,
		RecompensaEsperada: pirata.Recompensa,
		Detalle:        "El pirata ha sido entregado correctamente.",
	}

	if pirata.Recompensa > 200000000 {
		if rand.Float64() < 0.30 {
			resultado.EstadoFinal = pb.EstadoPirata_PERDIDO
			resultado.RecompensaEsperada = 0
			resultado.Detalle = "¡Los mercenarios han interceptado el transporte antes de la entrega!"
		}
	}

	return resultado, nil
}


func main() {
	listen, err := net.Listen("tcp", ":50053") 
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()

	servidor := &servidorSubmundo{}

	pb.RegisterSubmundoServer(grpcServer, servidor)

	log.Println("Servidor Submundo escuchando en el puerto 50053")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
