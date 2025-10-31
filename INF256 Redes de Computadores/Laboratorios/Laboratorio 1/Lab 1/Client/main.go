package main

import (
	"fmt"
	"net"
	"strings"
	"strconv"
)

func main() {

	// Conexión UDP para iniciar la partida
	udpAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	udpConn, _ := net.DialUDP("udp", nil, udpAddr)
	defer udpConn.Close()

	// Enviar mensaje de inicio
	udpConn.Write([]byte("Iniciar partida"))

	// Recibir respuesta del servidor
	buf := make([]byte, 1024)
	n, _, _ := udpConn.ReadFromUDP(buf)

	response := string(buf[:n])
	fmt.Println("Respuesta del servidor:", response)

	// Parsear la respuesta
	parts := strings.Split(response, ",")
	tcpPort := parts[1]
	numQuestions := parts[2]

	// Conexión TCP para responder preguntas
	tcpConn, _ := net.Dial("tcp", "localhost"+tcpPort)
	defer tcpConn.Close()

	// Recibir y responder preguntas
	num, _ := strconv.Atoi(numQuestions)
	for i := 0; i < num; i++ {
		// Recibir pregunta
		n, _ := tcpConn.Read(buf)
		question := string(buf[:n])
		fmt.Println("Pregunta:", question)

		// Enviar respuesta (simple simulación)
		var answer string
		fmt.Print("Tu respuesta: ")
		fmt.Scanln(&answer)
		tcpConn.Write([]byte(answer))

	}

	// Recibir resultados finales
	n, _ = tcpConn.Read(buf)
	result := string(buf[:n])
	fmt.Println("Resultado final:", result)
}

