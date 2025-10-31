package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	// Import your generated protobuf package here, adjust the path as needed
	pb "cdp/proto/grpc-server/proto"
)

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

// Desencripta un string en base64 cifrado con AES-256 CBC
func descifrarAES256(cifradoBase64 string, key []byte) ([]byte, error) {
	datos, err := base64.StdEncoding.DecodeString(cifradoBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(datos) < aes.BlockSize {
		return nil, fmt.Errorf("mensaje muy corto")
	}
	iv := datos[:aes.BlockSize]
	ciphertext := datos[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("tamaño de mensaje inválido")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Quitar padding PKCS7
	padding := int(ciphertext[len(ciphertext)-1])

	if padding > aes.BlockSize || padding == 0 {
		return nil, fmt.Errorf("padding inválido")
	}

	return ciphertext[:len(ciphertext)-padding], nil
}

func guardarResultadoEnArchivo(resultado MensajeResultado) error {
	const archivo = "resultados.json"
	var resultados []MensajeResultado

	data, err := os.ReadFile(archivo)
	if err == nil {
		if err := json.Unmarshal(data, &resultados); err != nil {
			log.Println("⚠️ Error parseando JSON previo, se sobrescribirá.")
			resultados = []MensajeResultado{} // reset
		}
	}

	resultados = append(resultados, resultado)

	jsonData, err := json.MarshalIndent(resultados, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	err = os.WriteFile(archivo, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	log.Println("💾 Resultado guardado exitosamente")
	return nil
}

func esMensajeValido(r MensajeResultado) (bool, string) {
	if r.IdGanador != r.IdEntrenador1 && r.IdGanador != r.IdEntrenador2 {
		return false, "el ID del ganador no coincide con los participantes"
	}

	// Validación muy básica de fecha: "YYYY-MM-DD"
	if len(r.Fecha) != 10 || r.Fecha[4] != '-' || r.Fecha[7] != '-' {
		return false, "la fecha no tiene formato YYYY-MM-DD"
	}

	return true, ""
}

func esDuplicado(resultado MensajeResultado, anteriores []MensajeResultado) bool {
	for _, r := range anteriores {
		if r.TorneoID == resultado.TorneoID &&
			r.IdEntrenador1 == resultado.IdEntrenador1 &&
			r.IdEntrenador2 == resultado.IdEntrenador2 &&
			r.Fecha == resultado.Fecha {
			return true
		}
	}
	return false
}

func entrenadoresValidos(client pb.LCPServiceClient, r MensajeResultado) bool {
	log.Printf("📡 Solicitando validación a LCP para entrenadores: %d (%s) y %d (%s)",
		r.IdEntrenador1, r.NombreEntrenador1,
		r.IdEntrenador2, r.NombreEntrenador2,
	)

	req := &pb.ValidacionRequest{
		IdEntrenador1: r.IdEntrenador1,
		IdEntrenador2: r.IdEntrenador2,
	}

	resp, err := client.ValidarEntrenadores(context.Background(), req)
	if err != nil {
		log.Printf("❌ Error en validación gRPC: %v", err)
		return false
	}
	if !resp.Valido {
		log.Printf("⚠️ Validación fallida: %s", resp.Mensaje)
	} else {
		log.Printf("✅ Validación exitosa para E1=%d y E2=%d", r.IdEntrenador1, r.IdEntrenador2)
	}

	return resp.Valido
}

func enviarResultadoValidado(ch *amqp.Channel, resultado MensajeResultado) error {
	data, err := json.Marshal(resultado)
	if err != nil {
		return fmt.Errorf("error serializando resultado validado: %w", err)
	}

	return ch.Publish(
		"",                     // exchange
		"resultados_validados", // routing key (nombre de la cola)
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("❌ Error conectando a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Error abriendo canal: %v", err)
	}
	defer ch.Close()

	// gRPC con LCP
	connLCP, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Error conectando a LCP: %v", err)
	}
	defer connLCP.Close()

	lcpClient := pb.NewLCPServiceClient(connLCP)

	_, err = ch.QueueDeclare(
		"resultados_combate", // nombre
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("❌ Error declarando cola: %v", err)
	}

	_, err = ch.QueueDeclare(
		"resultados_validados", // nombre
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Error declarando cola de resultados_validados: %v", err)
	}

	msgs, err := ch.Consume(
		"resultados_combate", // queue
		"",                   // consumer
		true,                 // auto-ack
		false, false, false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir cola: %v", err)
	}

	key := []byte("clave-de-32-bytes-12345678901234")
	if len(key) != 32 {
		log.Fatalf("❌ CLAVE_AES debe tener 32 bytes, actual: %d", len(key))
	}

	log.Println("📥 CDP escuchando mensajes en 'resultados_combate'...")

	for msg := range msgs {
		log.Println("📥 Mensaje recibido, intentando desencriptar...")
		jsonDesencriptado, err := descifrarAES256(string(msg.Body), key)
		if err != nil {
			log.Printf("❌ Error al desencriptar mensaje: %v\n", err)
			continue
		}

		var resultado MensajeResultado
		err = json.Unmarshal(jsonDesencriptado, &resultado)
		if err != nil {
			log.Printf("❌ Error al parsear JSON: %v\n", err)
			continue
		}

		ok, motivo := esMensajeValido(resultado)
		if !ok {
			log.Printf("❌ Resultado inválido: %s. Descarta.\n", motivo)
			continue
		}

		// Leer resultados anteriores
		var anteriores []MensajeResultado
		if data, err := os.ReadFile("resultados.json"); err == nil {
			_ = json.Unmarshal(data, &anteriores)
		}

		if esDuplicado(resultado, anteriores) {
			log.Println("⚠️ Resultado duplicado detectado. Descarta.")
			continue
		}

		if !entrenadoresValidos(lcpClient, resultado) {
			log.Println("❌ Entrenadores no válidos según LCP. Descarta.")
			continue
		}

		fmt.Printf("✅ Resultado recibido:\n")
		fmt.Printf("🏆 Torneo %d | 🗓️ %s\n", resultado.TorneoID, resultado.Fecha)
		fmt.Printf("👤 %s vs 👤 %s\n", resultado.NombreEntrenador1, resultado.NombreEntrenador2)
		fmt.Printf("🥇 Ganador: %s\n", resultado.NombreGanador)
		fmt.Printf("🥈 Perdedor: %s\n", resultado.NombrePerdedor)
		fmt.Println("------------------------------------------------")

		if err := guardarResultadoEnArchivo(resultado); err != nil {
			log.Printf("❌ Error al guardar en archivo: %v\n", err)
		}
		if err := enviarResultadoValidado(ch, resultado); err != nil {
			log.Printf("❌ Error enviando resultado validado a LCP: %v", err)
		} else {
			log.Println("📤 Resultado validado enviado a LCP vía RabbitMQ")
		}
	}

}
