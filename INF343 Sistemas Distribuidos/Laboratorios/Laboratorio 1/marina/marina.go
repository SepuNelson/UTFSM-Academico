package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "marina/proto/grpc-server/proto"
)

type servidorMarina struct {
	pb.UnimplementedMarinaServer
	piratasRecibidos map[string]*pb.Pirata
	ventasSubmundo   map[string]int
	mutex            sync.Mutex
}

func (s *servidorMarina) RecibirPirata(ctx context.Context, entrega *pb.EntregaPirata) (*pb.ResultadoEntrega, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	pirata := entrega.Pirata
	cazarrecompensas := entrega.NombreCazarrecompensas
	reputacion := entrega.Reputacion

	// Verificar si el pirata ya fue entregado (estado distinto a BUSCADO)
	if pirata.Estado != pb.EstadoPirata_BUSCADO {
		log.Printf("El pirata %s ya fue entregado o capturado previamente.", pirata.Nombre)
		return &pb.ResultadoEntrega{
			Aceptado:          false,
			RecompensaPagada: 0,
		}, nil
	}

	// Verificar la reputación del cazarrecompensas
	if reputacion < 50 {
		log.Printf("La reputación de %s es baja. Rechazo de entrega.", cazarrecompensas)
		return &pb.ResultadoEntrega{
			Aceptado:          false,
			RecompensaPagada: 0,
		}, nil
	}

	recompensa := pirata.Recompensa

	// Si el cazarrecompensas ha realizado ventas seguidas al Submundo, aplicar reducción en el pago
	if s.ventasSubmundo[cazarrecompensas] >= 3 {
		recompensa = recompensa / 2
		log.Printf("Pago reducido a %d Berries por ventas al Submundo", recompensa)
		// Notificar al Gobierno para actualizar la reputación del cazarrecompensas
	}

	// Incrementar el contador de ventas al Submundo si este es el destino
	if entrega.Metodo == pb.MetodoEntrega_SUBMUNDO {
		s.ventasSubmundo[cazarrecompensas]++
	}

	// Añadir el pirata a la lista de piratas recibidos
	s.piratasRecibidos[pirata.Id] = pirata

	log.Printf("El pirata %s ha sido aceptado. Pago de %d Berries a %s", pirata.Nombre, recompensa, cazarrecompensas)

	// Resultado de la entrega
	return &pb.ResultadoEntrega{
		Aceptado:          true,
		RecompensaPagada:  recompensa,
	}, nil
}

func (s *servidorMarina) RecibirAlerta(ctx context.Context, alerta *pb.AlertaActividad) (*emptypb.Empty, error) {
	log.Printf("Alerta de actividad detectada: %s (Nivel de riesgo: %d)", alerta.Mensaje, alerta.NivelRiesgo)
	// ??
	return &emptypb.Empty{}, nil
}

func (s *servidorMarina) ConsultarListaPiratas(ctx context.Context, _ *emptypb.Empty) (*pb.ListaPiratas, error) {
	var listaPiratas pb.ListaPiratas
	for _, pirata := range s.piratasRecibidos {
		listaPiratas.Piratas = append(listaPiratas.Piratas, pirata)
	}
	return &listaPiratas, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	servidor := &servidorMarina{
		piratasRecibidos: make(map[string]*pb.Pirata),
		ventasSubmundo:   make(map[string]int),
	}

	pb.RegisterMarinaServer(s, servidor)

	log.Println("Marina activa en puerto 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor de Marina: %v", err)
	}
}
