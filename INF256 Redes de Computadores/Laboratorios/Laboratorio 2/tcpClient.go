package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.188:8080")
	if err != nil {
        fmt.Println("Error :", err)
        return
    }
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error al recibir respuesta del servidor:", err)
        return
    }
    pregunta := string(buffer[:n])
    fmt.Printf("%s\n", pregunta)

	var respuesta string
	fmt.Print("Tu respuesta: ")
	fmt.Scanln(&(respuesta))

}