package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "lcp/proto/grpc-server/proto"

	"github.com/streadway/amqp"

	"google.golang.org/grpc"
)

/*
Variables and List Definitions:
- torneos: Lista de torneos disponibles.
- inscripciones: Lista de inscripciones a torneos.
- entrenadores: Lista de entrenadores.
*/

// Lista de Torneos
var torneos = []Tournament{
	{TorneoID: 001, Region: "Kanto"},
	{TorneoID: 002, Region: "Unova"},
	{TorneoID: 003, Region: "Sinnoh"},
	{TorneoID: 004, Region: "Johto"},
	{TorneoID: 005, Region: "Hoenn"},
}

var inscripciones = []*pb.Tournament_Inscription{}

/*
Structs Definitions:
- servidorLCP: Estructura del servidor gRPC.
- Tournament: Estructura para los torneos.
- Entrenador: Estructura para los entrenadores.
*/

// Estructura del servidor
type servidorLCP struct {
	pb.UnimplementedLCPServiceServer
	gymClient pb.GYMREGServiceClient
	trainers  []Entrenador
}

// Estructura de Torneos
type Tournament struct {
	TorneoID int32  `json:"torneo_id"`
	Region   string `json:"region"`
}

// Estructura para los entrenadores
type Entrenador struct {
	ID         string `json:"id"`
	Nombre     string `json:"nombre"`
	Region     string `json:"region"`
	Ranking    int    `json:"ranking"`
	Estado     string `json:"estado"`
	Suspension int    `json:"suspension"`
}

// RegistroParticipante define la información de participación en un torneo.
type RegistroParticipante struct {
	IdEntrenador   int32  `json:"id_entrenador"`
	Nombre         string `json:"nombre"`
	TorneoID       int32  `json:"torneo_id"`
	Resultado      string `json:"resultado"`
	RankingAntes   int32  `json:"ranking_antes"`
	RankingDespues int32  `json:"ranking_despues"`
}

type MensajeResultado struct {
	TorneoID          int32  `json:"torneo_id"`
	IdEntrenador1     int32  `json:"id_entrenador_1"`
	NombreEntrenador1 string `json:"nombre_entrenador_1"`
	IdEntrenador2     int32  `json:"id_entrenador_2"`
	NombreEntrenador2 string `json:"nombre_entrenador_2"`
	IdGanador         int32  `json:"id_ganador"`
	NombreGanador     string `json:"nombre_ganador"`
	IdPerdedor        int32  `json:"id_perdedor"`     // Nuevo
	NombrePerdedor    string `json:"nombre_perdedor"` // Nuevo
	Fecha             string `json:"fecha"`
	TipoMensaje       string `json:"tipo_mensaje"`
}

/*
Local Functions Definitions:
- cargarEntrenadores: Carga la lista de entrenadores desde un archivo JSON.
- inscripcionEntrenadores: Inscribe a los entrenadores en torneos aleatorios.
*/

// Método para recibir la lista de entrenadores
func cargarEntrenadores(filename string) ([]Entrenador, error) {

	// Leer el archivo JSON
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Leer el contenido del archivo
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Deserializar el contenido JSON en una lista de entrenadores
	var entrenadores []Entrenador
	if err := json.Unmarshal(bytes, &entrenadores); err != nil {
		return nil, err
	}
	// Imprimir la lista de entrenadores
	fmt.Println("Entrenadores Cargados Correctamente")
	return entrenadores, nil
}

// Método para agregar un nuevo entrenador a la lista
func agregarEntrenador(entrenadores []Entrenador, entrenador Entrenador) []Entrenador {

	// Verificar si el entrenador ya existe
	for _, e := range entrenadores {
		if e.ID == entrenador.ID {
			fmt.Printf("El entrenador con ID %s ya existe.\n", entrenador.ID)
			return entrenadores
		}
	}

	// Agregar el nuevo entrenador a la lista
	entrenadores = append(entrenadores, entrenador)
	fmt.Printf("Entrenador %s agregado exitosamente.\n", entrenador.Nombre)
	return entrenadores
}

// Inscribir entrenadores en torneos aleatorios
func inscripcionEntrenadores(entrenadores []Entrenador, ) {
	// Genera un número aleatorio para seleccionar un torneo
	rand.Seed(time.Now().UnixNano())

	// Para cada entrenador, se inscribe solo si su estado es "ACTIVO"
	for i, e := range entrenadores {

		if e.Estado == "Activo" {
			// Cambiar el estado a "INSCRITO" para que no se inscriba de nuevo
			entrenadores[i].Estado = "Inscrito"

			// Seleccionar un torneo aleatorio
			torneo := torneos[rand.Intn(len(torneos))]
			idTrainer, err := strconv.Atoi(e.ID)
			if err != nil {
				fmt.Printf("Error al convertir la ID del entrenador %s: %v\n", e.ID, err)
				continue
			}

			inscripcion := &pb.Tournament_Inscription{
				IdTournament: torneo.TorneoID,
				IdTrainer:    int32(idTrainer),
			}
			inscripciones = append(inscripciones, inscripcion)

			// Imprimir la inscripción
			fmt.Printf("Entrenador %s inscrito en el torneo %d de la región %s.\n", e.Nombre, torneo.TorneoID, torneo.Region)

		} else if e.Estado == "Expulsado" {

			// Si el entrenador está expulsado, no se inscribe
			fmt.Printf("Entrenador %s (ID %s) está expulsado y no puede inscribirse.\n", e.Nombre, e.ID)

		} else if e.Estado == "Suspendido" {

			// Bajar los números de la suspensión
			entrenadores[i].Suspension--
			if entrenadores[i].Suspension <= 0 {
				entrenadores[i].Estado = "Activo"
				entrenadores[i].Suspension = 0
			}

			// Si el entrenador está suspendido, no se inscribe
			fmt.Printf("Entrenador %s (ID %s) está suspendido y no puede inscribirse.\n", e.Nombre, e.ID)
		}
	}
}

// Obtener el entrenador por ID
func obtenerEntrenadorPorID(id int32, trainers []Entrenador) *pb.Trainer {
	for _, e := range trainers {
		idInt, _ := strconv.Atoi(e.ID)

		var estado pb.EstadoEntrenador
		switch e.Estado {
		case "Activo":
			estado = pb.EstadoEntrenador_ACTIVO
		case "Inscrito":
			estado = pb.EstadoEntrenador_INSCRITO
		case "Suspendido":
			estado = pb.EstadoEntrenador_SUSPENDIDO
		case "Expulsado":
			estado = pb.EstadoEntrenador_EXPULSADO
		default:
			estado = pb.EstadoEntrenador_ACTIVO
		}

		if int32(idInt) == id {
			return &pb.Trainer{
				Id:         int32(idInt),
				Nombre:     e.Nombre,
				Region:     e.Region,
				Ranking:    int32(e.Ranking),
				Estado:     estado,
				Suspension: int32(e.Suspension),
			}
		}
	}
	return nil
}

// Función auxiliar para obtener la región a partir del ID del torneo
func obtenerRegionPorTorneoID(torneoID int32) string {
	for _, t := range torneos {
		if t.TorneoID == torneoID {
			return t.Region
		}
	}
	return ""
}

// Función para asignar combates a los entrenadores
func combateEntrenadores(gymClient pb.GYMREGServiceClient, trainers []Entrenador) {

	// Duplicar la lista de inscripciones
	inscripcionesCopia := make([]*pb.Tournament_Inscription, len(inscripciones))
	copy(inscripcionesCopia, inscripciones)

	// Agrupar inscripciones por torneo
	porTorneo := make(map[int32][]*pb.Tournament_Inscription)
	for _, insc := range inscripcionesCopia {
		porTorneo[insc.IdTournament] = append(porTorneo[insc.IdTournament], insc)
	}

	combateID := int32(1)
	// Para cada grupo de inscripciones (por torneo)
	for torneoID, grupo := range porTorneo {

		// Asignar combates por pares
		for i := 0; i < len(grupo)-1; i += 2 {

			insc1 := grupo[i]
			insc2 := grupo[i+1]

			// Mostrar los entrenadores
			fmt.Printf("Asignación de combate de %s vs %s en %s torneo %d\n",
				obtenerEntrenadorPorID(insc1.IdTrainer, trainers).Nombre,
				obtenerEntrenadorPorID(insc2.IdTrainer, trainers).Nombre,
				obtenerRegionPorTorneoID(torneoID),
				torneoID)

			combate := &pb.Combate{
				CombateId:   combateID,
				TorneoId:    torneoID,
				Entrenador1: obtenerEntrenadorPorID(insc1.IdTrainer, trainers),
				Entrenador2: obtenerEntrenadorPorID(insc2.IdTrainer, trainers),
				Region:      obtenerRegionPorTorneoID(torneoID),
			}

			// Llamar a la función AsignarCombates()
			resp, err := gymClient.AsignarCombate(context.Background(), combate)
			if err != nil {
				fmt.Printf("Error asignando combate %d: %v\n", combateID, err)
			} else {
				fmt.Println(resp.Mensaje)
			}
			combateID++
		}
		// Si hay un número impar, el último queda sin rival
		if len(grupo)%2 != 0 {
			inscSolo := grupo[len(grupo)-1]
			entrenadorSolo := obtenerEntrenadorPorID(inscSolo.IdTrainer, trainers)
			fmt.Printf("Entrenador %s (ID %d) en torneo %d queda sin rival.\n", entrenadorSolo.Nombre, entrenadorSolo.Id, torneoID)
		}
	}
}

// Función para registrar las participaciones de los entrenadores en los torneos
func registroParticipaciones(torneoID int32, entrenador *pb.Trainer) RegistroParticipante {

	registro := RegistroParticipante{
		IdEntrenador:   entrenador.Id,
		Nombre:         entrenador.Nombre,
		TorneoID:       torneoID,
		Resultado:      "Actualizar",
		RankingAntes:   entrenador.Ranking,
		RankingDespues: 0,
	}

	return registro
}

/*
gRPC Functions Definitions:
- PublicarTorneos: Método para publicar torneos.
- InscribirseTorneo: Método para inscribirse en un torneo.
*/

// PublicarTorneos: Método para publicar torneos
func (s *servidorLCP) PublicarTorneos(ctx context.Context, req *pb.Empty) (*pb.Tournament_List, error) {

	// Crear una lista vacía para los torneos en formato protobuf
	var torneosProto []*pb.Tournament

	// Itera sobre la lista de torneos y los convierte a formato protobuf
	for _, t := range torneos {
		torneosProto = append(torneosProto, &pb.Tournament{
			Id:     t.TorneoID,
			Region: t.Region,
		})
	}

	// Imprime la lista de torneos
	return &pb.Tournament_List{Torneos: torneosProto}, nil
}

// InscribirseTorneo: Método para inscribirse en un torneo
func (s *servidorLCP) InscribirseTorneo(ctx context.Context, req *pb.InscripcionRequest) (*pb.Response, error) {

	// Datos de inscripción
	entrenador := req.GetEntrenador()
	idTorneo := req.GetIdTorneo()

	// Registrar al entrenador cliente
	registroParticipaciones(idTorneo, entrenador)

	entrenadorLocal := Entrenador{
		ID:         strconv.Itoa(int(entrenador.GetId())),
		Nombre:     entrenador.GetNombre(),
		Region:     entrenador.GetRegion(),
		Ranking:    int(entrenador.GetRanking()),
		Estado:     entrenador.GetEstado().String(),
		Suspension: int(entrenador.GetSuspension()),
	}

	s.trainers = agregarEntrenador(s.trainers, entrenadorLocal)

	// ACTIVO
	if entrenador.Estado == pb.EstadoEntrenador_ACTIVO {

		// Verifica si ya está inscrito en cualquier torneo
		for _, insc := range inscripciones {
			if insc.GetIdTrainer() == entrenador.GetId() {
				return &pb.Response{Mensaje: "El entrenador ya está inscrito en un torneo y no puede inscribirse en otro.\n"}, nil
			}
		}

		// Cambia estado a "inscrito"
		entrenador.Estado = pb.EstadoEntrenador_INSCRITO

		// Agrega la inscripción
		inscripcion := &pb.Tournament_Inscription{
			IdTournament: idTorneo,
			IdTrainer:    entrenador.GetId(),
		}
		inscripciones = append(inscripciones, inscripcion)

		time.Sleep(2 * time.Second)
		// Inscribir entrenadores en torneos
		inscripcionEntrenadores(s.trainers)
		time.Sleep(3 * time.Second)

		// Después de enviar la respuesta, dispara la asignación de combates en background
		go func() {
			combateEntrenadores(s.gymClient, s.trainers)
		}()

		// Inscripción exitosa
		return &pb.Response{Mensaje: "Inscripción exitosa.\n"}, nil

		// EXPULSADO
	} else if entrenador.Estado == pb.EstadoEntrenador_EXPULSADO {

		// Inscripción rechazada
		return &pb.Response{Mensaje: "Entrenador expulsado, no puede inscribirse.\n"}, nil

		// SUSPENDIDO
	} else if entrenador.Estado == pb.EstadoEntrenador_SUSPENDIDO {

		// Bajar los números de la suspensión
		entrenador.Suspension--
		if entrenador.Suspension <= 0 {
			entrenador.Estado = pb.EstadoEntrenador_ACTIVO
			entrenador.Suspension = 0
		}

		// Inscripción rechazada
		return &pb.Response{Mensaje: "Entrenador suspendido, por ahora no puede inscribirse.\n"}, nil

	}
	// Si el estado no es válido
	return &pb.Response{Mensaje: "Estado de entrenador no válido para inscripción.\n"}, nil
}

func (s *servidorLCP) ValidarEntrenadores(ctx context.Context, req *pb.ValidacionRequest) (*pb.ValidacionResponse, error) {
	log.Printf("📥 LCP recibió solicitud de validación: E1=%d, E2=%d", req.IdEntrenador1, req.IdEntrenador2)

	e1 := obtenerEntrenadorPorID(req.IdEntrenador1, s.trainers)
	e2 := obtenerEntrenadorPorID(req.IdEntrenador2, s.trainers)

	if e1 == nil || e2 == nil {
		log.Printf("❌ Uno o ambos entrenadores no existen. E1: %v | E2: %v", e1, e2)
		return &pb.ValidacionResponse{Valido: false, Mensaje: "Uno o ambos entrenadores no existen"}, nil
	}

	log.Printf("🔍 E1: ID=%d, Estado=%s, Nombre=%s", e1.Id, e1.Estado.String(), e1.Nombre)
	log.Printf("🔍 E2: ID=%d, Estado=%s, Nombre=%s", e2.Id, e2.Estado.String(), e2.Nombre)

	if e1.Estado != pb.EstadoEntrenador_INSCRITO || e2.Estado != pb.EstadoEntrenador_INSCRITO {
		log.Println("🚫 Uno o ambos entrenadores no están activos.")
		return &pb.ValidacionResponse{Valido: false, Mensaje: "Uno o ambos entrenadores no están activos"}, nil
	}

	log.Println("✅ Entrenadores válidos y activos.")
	return &pb.ValidacionResponse{Valido: true, Mensaje: "Validación exitosa"}, nil
}

func (s *servidorLCP) RegistrarParticipacion(ctx context.Context, inscripcion *pb.InscripcionRequest) (*pb.RegistroParticipante, error) {

	// Obtener el entrenador del request
	entrenador := inscripcion.GetEntrenador()
	idtorneo := inscripcion.GetIdTorneo()

	fmt.Print("%s", entrenador.Estado)

	if entrenador.Estado == pb.EstadoEntrenador_ACTIVO {
		registro := registroParticipaciones(idtorneo, entrenador)

		// Convertir RegistroParticipante a pb.RegistroParticipante
		registroProto := &pb.RegistroParticipante{
			IdEntrenador:   registro.IdEntrenador,
			Nombre:         registro.Nombre,
			TorneoID:       registro.TorneoID,
			Resultado:      registro.Resultado,
			RankingAntes:   registro.RankingAntes,
			RankingDespues: registro.RankingDespues,
		}

		// Respuesta exitosa
		return registroProto, nil
	} else {
		// Si el entrenador está expulsado, no se registra
		return nil, fmt.Errorf("el entrenador %s (ID %d) está expulsado y no puede registrarse", entrenador.Nombre, entrenador.Id)
	}
}

func main() {
	// Establecer conexión con el servidor gRPC LCP
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Establecer conexión con el servidor gRPC GYMREG
	connGym, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con GYMREG: %v", err)
	}
	defer connGym.Close()

	// Crear el cliente gRPC GYMREG
	gymClient := pb.NewGYMREGServiceClient(connGym)

	// Leer y cargar el archivo JSON
	trainers, err := cargarEntrenadores("entrenadores.json")
	if err != nil {
		log.Fatalf("Error cargando entrenadores: %v", err)
	}

	// Crear el servidor gRPC LCP

	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterLCPServiceServer(grpcServer, &servidorLCP{
			gymClient: gymClient,
			trainers:  trainers,
		})

		fmt.Println("✅ LCP Server is running on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	go func() {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatalf("❌ Error conectando a RabbitMQ desde LCP: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("❌ Error abriendo canal RabbitMQ en LCP: %v", err)
		}
		defer ch.Close()

		_, err = ch.QueueDeclare("resultados_validados", false, false, false, false, nil)
		if err != nil {
			log.Fatalf("❌ Error declarando cola en LCP: %v", err)
		}

		msgs, err := ch.Consume("resultados_validados", "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("❌ Error suscribiéndose a resultados_validados: %v", err)
		}

		log.Println("📡 LCP escuchando resultados oficiales desde CDP...")

		for msg := range msgs {
			var r MensajeResultado
			if err := json.Unmarshal(msg.Body, &r); err != nil {
				log.Printf("❌ Error parseando resultado recibido: %v", err)
				continue
			}

			fmt.Println("📣 Resultado oficial recibido (CDP ®):")
			fmt.Printf("🏆 Torneo %d | 🗓️ %s\n", r.TorneoID, r.Fecha)
			fmt.Printf("👤 %s vs 👤 %s\n", r.NombreEntrenador1, r.NombreEntrenador2)
			fmt.Printf("🥇 Ganador: %s\n", r.NombreGanador)
			fmt.Printf("🥈 Perdedor: %s\n", r.NombrePerdedor)
			fmt.Println("------------------------------------------------")

			var ganador, perdedor *Entrenador
			for i := range trainers {
				idInt, _ := strconv.Atoi(trainers[i].ID)
				if int32(idInt) == r.IdGanador {
					ganador = &trainers[i]
				}
				if int32(idInt) == r.IdPerdedor {
					perdedor = &trainers[i]
				}
			}

			if ganador != nil && perdedor != nil {
				fmt.Printf("Ranking antes - Ganador: %s (%d), Perdedor: %s (%d)\n",
					ganador.Nombre, ganador.Ranking, perdedor.Nombre, perdedor.Ranking)

				ganador.Ranking += 30
				perdedor.Ranking -= 30

				fmt.Printf("Ranking después - Ganador: %s (%d), Perdedor: %s (%d)\n",
					ganador.Nombre, ganador.Ranking, perdedor.Nombre, perdedor.Ranking)

				// Notificar al SNP el cambio de ranking (ganador)
				notifGanador := map[string]interface{}{
					"tipo_mensaje":      "ranking_actualizado",
					"id_entrenador":     ganador.ID,
					"nombre_entrenador": ganador.Nombre,
					"nuevo_ranking":     ganador.Ranking,
					"fecha":             time.Now().Format("2006-01-02"),
				}
				body, _ := json.Marshal(notifGanador)
				ch.Publish("", "notificaciones_snp", false, false, amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})

				// Notificar al SNP el cambio de ranking (perdedor)
				notifPerdedor := map[string]interface{}{
					"tipo_mensaje":      "ranking_actualizado",
					"id_entrenador":     perdedor.ID,
					"nombre_entrenador": perdedor.Nombre,
					"nuevo_ranking":     perdedor.Ranking,
					"fecha":             time.Now().Format("2006-01-02"),
				}
				body2, _ := json.Marshal(notifPerdedor)
				ch.Publish("", "notificaciones_snp", false, false, amqp.Publishing{
					ContentType: "application/json",
					Body:        body2,
				})

			} else {
				fmt.Println("No se encontró al ganador o perdedor en la lista de entrenadores.")
			}

			nuevoTorneo := Tournament{
				TorneoID: int32(len(torneos) + 1),
				Region:   "Kalos", // HARDCODEADO NO SE COMO HACERLO AQUII
			}
			torneos = append(torneos, nuevoTorneo)

			notifTorneo := map[string]interface{}{
				"tipo_mensaje": "nuevo_torneo",
				"torneo_id":    nuevoTorneo.TorneoID,
				"region":       nuevoTorneo.Region,
				"fecha":        time.Now().Format("2006-01-02"),
			}
			bodyTorneo, _ := json.Marshal(notifTorneo)
			ch.Publish("", "notificaciones_snp", false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        bodyTorneo,
			})

		}
	}()
	select {}
}
