package main

import (
	"fmt"
	"net"
)

func main() {
	servidor := "192.168.1.188"

	conn, err := net.Dial("tcp", servidor)
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
    respuesta := string(buffer[:n])
    fmt.Printf("Respuesta del servidor: %s\n", respuesta)
}
