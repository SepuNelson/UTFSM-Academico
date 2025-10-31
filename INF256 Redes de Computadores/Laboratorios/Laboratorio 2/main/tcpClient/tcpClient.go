package main

import (
    "fmt"
    "net"
	"strings"
)

func main() {
	// Extraer la IP y el puerto de la respuesta

    // Conectarse al servidor TCP
    tcpConn, err := net.Dial("tcp", "192.168.1.188:8080")
    if err != nil {
        fmt.Println("Error al conectar con el servidor TCP:", err)
        return
    }
    defer tcpConn.Close()


	// Comunicacion cliente-servidor
    for {
        buffer := make([]byte, 1024)
        n, err := tcpConn.Read(buffer)
        if err != nil {
            if err.Error() == "EOF" {
                fmt.Println("El servidor ha cerrado la conexión.")
                break
            }
            fmt.Println("Error al recibir pregunta del servidor:", err)
            break
        }
        pregunta := string(buffer[:n])
        if strings.HasPrefix(pregunta, "Resultados:") {
            fmt.Printf("%s\n", pregunta)
            break
        }
        fmt.Printf("Pregunta recibida: %s\n", pregunta)

        var respuesta string
        fmt.Print("Tu respuesta: ")
        fmt.Scanln(&(respuesta))

        //Enviar respuesta al servidor
        tcpConn.Write([]byte(respuesta))
    }

    fmt.Println("Conexión TCP cerrada.")
}
