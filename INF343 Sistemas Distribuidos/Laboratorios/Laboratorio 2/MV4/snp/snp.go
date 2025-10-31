package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

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

	_, err = ch.QueueDeclare("notificaciones_snp", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Error declarando cola notificaciones_snp: %v", err)
	}

	msgs, err := ch.Consume("notificaciones_snp", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Error suscribiéndose a notificaciones_snp: %v", err)
	}

	log.Println("📡 SNP escuchando notificaciones de LCP...")

	for msg := range msgs {
		var notif map[string]interface{}
		if err := json.Unmarshal(msg.Body, &notif); err != nil {
			log.Printf("❌ Error parseando notificación: %v", err)
			continue
		}

		log.Printf("📥 SNP recibió mensaje: %s", string(msg.Body))

		tipo, _ := notif["tipo_mensaje"].(string)
		switch tipo {
		case "ranking_actualizado":
			id, _ := notif["id_entrenador"].(string)
			queueName := "notificaciones_entrenador_" + id
			ch.QueueDeclare(queueName, false, false, false, false, nil)
			ch.Publish("", queueName, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Body,
			})
			log.Printf("📤 SNP reenvió notificación de ranking a %s", queueName)
		case "nuevo_torneo":
			ch.QueueDeclare("notificaciones_torneos", false, false, false, false, nil)
			ch.Publish("", "notificaciones_torneos", false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Body,
			})
			log.Println("📤 SNP reenvió notificación de nuevo torneo a notificaciones_torneos")
		default:
			log.Printf("⚠️ Tipo de mensaje no manejado: %s", tipo)
		}

		// Print para separar mensajes en consola
		log.Println("--------------------------------------------------")
		time.Sleep(500 * time.Millisecond)
	}
}
