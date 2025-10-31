package main

import (
	"context"
	"encoding/csv"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "gobierno/proto/grpc-server/proto" 
)

type servidorGobierno struct {
	pb.UnimplementedGobiernoMundialServer
	piratas        map[string]*pb.Pirata
	reputaciones   map[string]float32
	ventasSubmundo map[string]int
	mutex          sync.Mutex
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	servidor := &servidorGobierno{
		piratas:        make(map[string]*pb.Pirata),
		reputaciones:   make(map[string]float32),
		ventasSubmundo: make(map[string]int),
	}

	err = servidor.cargarPiratasDesdeCSV("piratas.csv")
	if err != nil {
		log.Fatalf("Error al cargar CSV: %v", err)
	}

	pb.RegisterGobiernoMundialServer(grpcServer, servidor)

	log.Println("Gobierno Mundial activo en el puerto 50051 🚀")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func (s *servidorGobierno) cargarPiratasDesdeCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines[1:] { // saltar encabezado
		recompensa, _ := strconv.ParseUint(strings.TrimSpace(line[2]), 10, 64)

		var peligrosidad pb.Peligrosidad
		switch strings.ToLower(strings.TrimSpace(line[3])) {
		case "bajo":
			peligrosidad = pb.Peligrosidad_BAJA
		case "medio", "media":
			peligrosidad = pb.Peligrosidad_MEDIA
		case "alto":
			peligrosidad = pb.Peligrosidad_ALTA
		default:
			log.Printf("Peligrosidad desconocida '%s', usando BAJA por defecto\n", line[3])
			peligrosidad = pb.Peligrosidad_BAJA
		}

		estadoTexto := strings.ToUpper(strings.TrimSpace(line[4]))
		estadoEnum, ok := pb.EstadoPirata_value[estadoTexto]
		if !ok {
			log.Printf("Estado desconocido '%s', usando BUSCADO por defecto\n", line[4])
			estadoEnum = int32(pb.EstadoPirata_BUSCADO)
		}

		s.piratas[line[0]] = &pb.Pirata{
			Id:           strings.TrimSpace(line[0]),
			Nombre:       strings.TrimSpace(line[1]),
			Recompensa:   recompensa,
			Peligrosidad: peligrosidad,
			Estado:       pb.EstadoPirata(estadoEnum),
		}
	}

	log.Printf("Se cargaron %d piratas desde %s\n", len(s.piratas), path)
	return nil
}


func (s *servidorGobierno) ListarPiratasBuscados(_ context.Context, _ *emptypb.Empty) (*pb.ListaPiratas, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	lista := &pb.ListaPiratas{}
	for _, p := range s.piratas {
		if p.Estado == pb.EstadoPirata_BUSCADO {
			lista.Piratas = append(lista.Piratas, p)
		}
	}
	return lista, nil
}

func (s *servidorGobierno) AlertarMarina(_ context.Context, alerta *pb.AlertaActividad) (*emptypb.Empty, error) {
	log.Printf("Alerta a la Marina: %s (Nivel de riesgo: %d)\n", alerta.Mensaje, alerta.NivelRiesgo)
	return &emptypb.Empty{}, nil
}


func (s *servidorGobierno) RegistrarCaptura(_ context.Context, req *pb.SolicitudCaptura) (*pb.ResultadoCaptura, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	pirata, ok := s.piratas[req.IdPirata]
	if !ok {
		return &pb.ResultadoCaptura{
			Exito:   false,
			Mensaje: "Pirata no encontrado",
		}, nil
	}

	if pirata.Estado != pb.EstadoPirata_BUSCADO {
		return &pb.ResultadoCaptura{
			Exito:   false,
			Mensaje: "Pirata ya fue capturado",
		}, nil
	}

	pirata.Estado = pb.EstadoPirata_CAPTURADO

	// Si ya llevaba 3 ventas seguidas en el submundo, avisar a la Marina
	if s.ventasSubmundo[req.NombreCazarrecompensas] >= 3 {
		log.Printf("Alerta: %s ha vendido 3 veces seguidas en el Submundo. Se notificará a la Marina.\n", req.NombreCazarrecompensas)
	}

	return &pb.ResultadoCaptura{
		Exito:   true,
		Mensaje: fmt.Sprintf("Pirata %s capturado exitosamente", pirata.Nombre),
	}, nil
}

func (s *servidorGobierno) ActualizarReputacion(_ context.Context, estado *pb.EstadoCazarrecompensas) (*emptypb.Empty, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	reputacion := estado.Reputacion
	nueva := reputacion * 0.95
	if nueva < 0 {
		nueva = 0
	}

	s.reputaciones[estado.NombreCazarrecompensas] = nueva
	log.Printf("Reputación de %s actualizada a %.2f\n", estado.NombreCazarrecompensas, nueva)
	return &emptypb.Empty{}, nil
}

func (s *servidorGobierno) RegistrarCazarrecompensas(_ context.Context, estado *pb.EstadoCazarrecompensas) (*pb.ResultadoCaptura, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	nombre := estado.NombreCazarrecompensas
	if _, existe := s.reputaciones[nombre]; existe {
		return &pb.ResultadoCaptura{
			Exito:   false,
			Mensaje: fmt.Sprintf("El nombre '%s' ya está registrado.", nombre),
		}, nil
	}

	s.reputaciones[nombre] = 100.0
	log.Printf("Cazarrecompensas '%s' registrado con reputación inicial de 100.", nombre)

	return &pb.ResultadoCaptura{
		Exito:   true,
		Mensaje: "Registro exitoso.",
	}, nil
}

func (s *servidorGobierno) ConsultarEstadoCazarrecompensas(_ context.Context, req *pb.EstadoCazarrecompensas) (*pb.EstadoCazarrecompensas, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	reputacion, existe := s.reputaciones[req.NombreCazarrecompensas]
	if !existe {
		log.Printf("Consulta fallida: %s no está registrado.", req.NombreCazarrecompensas)
		return &pb.EstadoCazarrecompensas{
			NombreCazarrecompensas: req.NombreCazarrecompensas,
			Reputacion:             0,
			TotalGanado:            0,
			Historial:              []string{},
		}, nil
	}

	log.Printf("Consulta de estado: %s tiene reputación %.2f", req.NombreCazarrecompensas, reputacion)

	return &pb.EstadoCazarrecompensas{
		NombreCazarrecompensas: req.NombreCazarrecompensas,
		Reputacion:             reputacion,
	}, nil
}

func (s *servidorGobierno) ActualizarEstadoPirata(_ context.Context, pirata *pb.Pirata) (*emptypb.Empty, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p, ok := s.piratas[pirata.Id]
	if !ok {
		log.Printf("Pirata con ID %s no encontrado para actualizar estado.", pirata.Id)
		return &emptypb.Empty{}, nil
	}

	p.Estado = pirata.Estado
	log.Printf("Estado del pirata '%s' actualizado a %s", p.Nombre, p.Estado.String())

	return &emptypb.Empty{}, nil
}
