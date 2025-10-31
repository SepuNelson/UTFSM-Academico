package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	cryptorand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	pb "gym_reg/proto/grpc-server/proto"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type servidorGYMREG struct {
	pb.UnimplementedGYMREGServiceServer
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

func ejecutarCombates(entrenador1, entrenador2 *pb.Trainer) *pb.Trainer {
	time.Sleep(5 * time.Second)
	diff := float64(entrenador1.Ranking - entrenador2.Ranking)
	k := 100.0
	prob := 1.0 / (1.0 + math.Exp(-diff/k))
	if rand.Float64() <= prob {
		return entrenador1
	}
	return entrenador2
}

func cifrarAES256(msg []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := cryptorand.Read(iv); err != nil {
		return "", err
	}

	padding := aes.BlockSize - len(msg)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	msg = append(msg, padtext...)

	ciphertext := make([]byte, len(msg))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, msg)

	final := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(final), nil
}

func enviarRabbitMQ(mensaje string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Conexión Docker
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("resultados_combate", false, false, false, false, nil)
	if err != nil {
		return err
	}

	return ch.Publish("", "resultados_combate", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(mensaje),
	})
}

// Implementación del método gRPC con cambios clave
func (s *servidorGYMREG) AsignarCombate(ctx context.Context, req *pb.Combate) (*pb.Response, error) {
	log.Printf("📥 Entrenadores recibidos | Estado1: %s | Estado2: %s",
		req.Entrenador1.Estado.String(), req.Entrenador2.Estado.String())

	log.Printf("\n🔥 Combate recibido | ID: %d | Torneo: %d", req.CombateId, req.TorneoId)

	// Validación de estados (Inscritos)
	if req.Entrenador1.Estado != pb.EstadoEntrenador_INSCRITO || req.Entrenador2.Estado != pb.EstadoEntrenador_INSCRITO {
		log.Printf("⛔ Entrenadores no inscritos | ID1: %d | ID2: %d", req.Entrenador1.Id, req.Entrenador2.Id)
		return nil, fmt.Errorf("ambos entrenadores deben estar inscritos")
	}

	// Ejecutar combate
	entrenadorGanador := ejecutarCombates(req.Entrenador1, req.Entrenador2)
	log.Printf("⚔️ Combate ejecutado | Ganador: %s", entrenadorGanador.Nombre)

	var entrenadorPerdedor *pb.Trainer
	if entrenadorGanador.Id == req.Entrenador1.Id {
		entrenadorPerdedor = req.Entrenador2
	} else {
		entrenadorPerdedor = req.Entrenador1
	}

	// Actualizar estados
	req.Entrenador1.Estado = pb.EstadoEntrenador_ESPERANDO_RESULTADO
	req.Entrenador2.Estado = pb.EstadoEntrenador_ESPERANDO_RESULTADO
	log.Println("🕒 Estados actualizados: Esperando resultado")

	// Construir mensaje
	mensaje := MensajeResultado{
		TorneoID:          req.TorneoId,
		IdEntrenador1:     req.Entrenador1.Id,
		NombreEntrenador1: req.Entrenador1.Nombre,
		IdEntrenador2:     req.Entrenador2.Id,
		NombreEntrenador2: req.Entrenador2.Nombre,
		IdGanador:         entrenadorGanador.Id,
		NombreGanador:     entrenadorGanador.Nombre,
		IdPerdedor:        entrenadorPerdedor.Id,
		NombrePerdedor:    entrenadorPerdedor.Nombre,
		Fecha:             time.Now().Format("2006-01-02"),
		TipoMensaje:       "resultado_combate",
	}

	jsonBytes, err := json.Marshal(mensaje)
	if err != nil {
		log.Fatal("❌ Error serializando JSON:", err)
	}

	key := []byte("clave-de-32-bytes-12345678901234") // hardcoded for simplicity
	if len(key) != 32 {
		log.Fatal("❌ CLAVE_AES debe ser de 32 bytes (actual:", len(key), ")")
	}
	log.Printf("🔑 Clave AES cargada correctamente (%d bytes)", len(key))

	cifrado, err := cifrarAES256(jsonBytes, key)
	if err != nil {
		log.Fatal("❌ Error cifrando:", err)
	}
	log.Printf("🔐 Mensaje cifrado (base64): %s...", cifrado[:30])

	if err := enviarRabbitMQ(cifrado); err != nil {
		log.Fatal("❌ Error enviando a RabbitMQ:", err)
	}
	log.Println("📤 Mensaje enviado a cola 'resultados_combate'")

	return &pb.Response{Mensaje: "Combate procesado exitosamente"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("⛔ No se pudo escuchar en :50052: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGYMREGServiceServer(s, &servidorGYMREG{})
	log.Println("🏟️ Gimnasio Regional iniciado (puerto :50052)")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("⛔ Fallo al iniciar servidor gRPC: %v", err)
	}
}
