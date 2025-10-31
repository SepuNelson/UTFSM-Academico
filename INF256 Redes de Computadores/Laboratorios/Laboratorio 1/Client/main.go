package main

import (
    "fmt"
    "net"
    "strings"
)

func main() {
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Error resolviendo dirección UDP:", err)
        return
    }

    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("Error al conectar con el servidor UDP:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Conectado al servidor UDP")

    // Enviar solicitud de inicio de trivia
    conn.Write([]byte("Inicio Trivia"))

    // Recibir respuesta del servidor
    buffer := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        fmt.Println("Error al recibir respuesta del servidor:", err)
        return
    }
    respuesta := string(buffer[:n])
    fmt.Printf("Respuesta del servidor: %s\n", respuesta)

    // Extraer la IP y el puerto de la respuesta
    parts := strings.Split(respuesta, ":")
    if len(parts) != 3 {
        fmt.Println("Respuesta del servidor no válida")
        return
    }
    tcpAddr := fmt.Sprintf("%s:%s", parts[1], parts[2])

    // Conectarse al servidor TCP
    tcpConn, err := net.Dial("tcp", tcpAddr)
    if err != nil {
        fmt.Println("Error al conectar con el servidor TCP:", err)
        return
    }
    defer tcpConn.Close()

    fmt.Println("Conectado al servidor TCP")

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
            // Enviar mensaje de finalización al servidor
            tcpConn.Write([]byte("Fin Trivia"))
            break
        }
        fmt.Printf("Pregunta recibida: %s\n", pregunta)

        var respuesta string
        fmt.Print("Tu respuesta: ")
        fmt.Scanln(&(respuesta))

        respuesta = strings.Title(respuesta)
        //Enviar respuesta al servidor
        tcpConn.Write([]byte(respuesta))
    }

    fmt.Println("Conexión TCP cerrada.")
}

