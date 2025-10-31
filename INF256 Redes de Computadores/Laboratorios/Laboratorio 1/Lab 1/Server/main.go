package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"strings"
)

const (
	udpPort = ":8080"
	tcpPort = ":8081"
)

var questions = map[string]string{
	"¿Cuál es la capital de Chile?": "Santiago",
    "¿Cuál es el río más largo del mundo?": "Amazonas",
    "¿Cuál es el océano más grande?": "Pacífico",
    "¿Cuál es el país más grande del mundo?": "Rusia",
    "¿Cuál es el planeta más grande del sistema solar?": "Júpiter",
    "¿Cuál es el país más poblado del mundo?": "China",
    "¿Cuál es el país más pequeño del mundo?": "Vaticano",
    "¿Cuál es el país más poblado de América?": "Estados Unidos",
    "¿Cuál es el país más poblado de Europa?": "Rusia",
    "¿Cuál es el país más poblado de África?": "Nigeria",
}

// Manejar conexión UDP
func handleUDPConnection(conn *net.UDPConn, numQuestions int) {
	fmt.Println()
	defer conn.Close()
	buf := make([]byte, 1024)
	_, addr, _ := conn.ReadFromUDP(buf)
	fmt.Println("Solicitud recibida desde:", addr)

	// Enviar información al cliente (IP, puerto y número de preguntas)
	response := fmt.Sprintf("OK,%s,%d", tcpPort, numQuestions)
	conn.WriteToUDP([]byte(response), addr)
}

// Manejar conexión TCP
func handleTCPConnection(conn net.Conn, numQuestions int) {
	defer conn.Close()

	score := 0
	for i := 0; i < numQuestions; i++ {
		// Seleccionar una pregunta aleatoria
		rand.Seed(time.Now().UnixNano())
		for question, answer := range questions {

			delete(questions, question)

			// Enviar pregunta al cliente
			conn.Write([]byte(question))

			// Recibir respuesta del cliente
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)

			clientAnswer := strings.TrimSpace(string(buf[:n]))

			fmt.Println("Respuesta del cliente:", clientAnswer)

			// Verificar respuesta
			if clientAnswer == answer {
				score++
			}
			break
		}
	}

	// Enviar resultado al cliente
	result := fmt.Sprintf("Resultado: %d", score)
	conn.Write([]byte(result))

}

func main() {
	// Configurar conexión UDP
	udpAddr, _ := net.ResolveUDPAddr("udp", udpPort)
	udpConn, _ := net.ListenUDP("udp", udpAddr)

	// Generar número aleatorio de preguntas (entre 3 y 7)
	numQuestions := rand.Intn(5) + 3

	go handleUDPConnection(udpConn, numQuestions)

	// Configurar conexión TCP
	tcpListener, _ := net.Listen("tcp", tcpPort)
	defer tcpListener.Close()

	
	conn, _ := tcpListener.Accept()

	// Aquí podrías extraer el número de preguntas del mensaje UDP previo
	handleTCPConnection(conn, numQuestions) // Aquí 5 es un número de ejemplo, debe ser el que se envía desde UDP

}
