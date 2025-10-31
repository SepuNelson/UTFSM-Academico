package main

import (
    "fmt"
    "math/rand"
    "net"
    "time"
)

var preguntas = map[string]string{
    "¿Cuál es la capital de Chile?": "Santiago",
    "¿Cuál es el río más largo del mundo?": "Amazonas",
    "¿Cuál es el océano más grande?": "Pacífico",
    "¿Cuál es el país más grande del mundo?": "Rusia",
    "¿Cuál es el planeta más grande del sistema solar?": "Júpiter",
    "¿Cuál es el país más poblado del mundo?": "China",
    "¿Cuál es el país más pequeño del mundo?": "Vaticano",
    "¿Cuál es el país más poblado de Europa?": "Rusia",
    "¿Cuál es el país más poblado de África?": "Nigeria",
    "¿Como se llama el hueso más largo del cuerpo humano?": "Fémur",
    "¿Cuantos años hay en un lustro?": "5",
    "¿Cuantos años hay en una década?": "10",
    "¿Cuantos años hay en un siglo?": "100",
    "¿Cuantos años hay en un milenio?": "1000",
    "¿Cuantos días tiene un año bisiesto?": "366",
}

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("Error al resolver dirección UDP: ", err)
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error al escuchar en UDP: ", err)
        return
    }
    defer conn.Close()

    fmt.Println("Servidor UDP escuchando en el puerto 8080")

    handleClient(conn)
    
}

func handleClient(conn *net.UDPConn) {
    buffer := make([]byte, 1024)
    n, clientAddr, err := conn.ReadFromUDP(buffer)
    if err != nil {
        fmt.Println("Error al leer del cliente:", err)
        return
    }
    fmt.Printf("Mensaje recibido del cliente: %s\n", string(buffer[:n]))

    rand.Seed(time.Now().UnixNano())
    numPreguntas := rand.Intn(5) + 3

    respuesta := fmt.Sprintf("%d:localhost:9090", numPreguntas)
    _, err = conn.WriteToUDP([]byte(respuesta), clientAddr)
    if err != nil {
        fmt.Println("Error al enviar respuesta al cliente:", err)
        return
    }
    fmt.Printf("Enviado al cliente: %s\n", respuesta)

    startTCPServer(numPreguntas)

}

func startTCPServer(numPreguntas int) {
    ln, err := net.Listen("tcp", ":9090")
    if err != nil {
        fmt.Println("Error al iniciar el servidor TCP:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Servidor TCP escuchando en el puerto 9090")

    conn, err := ln.Accept()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Conexión TCP establecida")

    var respuestasCorrectas int
    var respuestasTotales int

    for i := 0; i < numPreguntas; i++ {

        // Seleccionar una pregunta aleatoria
		rand.Seed(time.Now().UnixNano())
        for pregunta, respuesta := range preguntas {

            conn.Write([]byte(pregunta))
            fmt.Println("Pregunta enviada:", pregunta)

            buffer := make([]byte, 1024)
            n, err := conn.Read(buffer)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            respuestaCliente := string(buffer[:n])
            fmt.Printf("Respuesta del cliente: %s\n", respuestaCliente)

            // Verificar si la respuesta es correcta
            if respuestaCliente == respuesta {
                fmt.Println("Respuesta correcta")
                respuestasCorrectas++
            } else {
                fmt.Println("Respuesta incorrecta")
            }
            respuestasTotales++

            delete(preguntas, pregunta)

            break
        }
    }

    // Enviar resultados finales al cliente
    resultados := fmt.Sprintf("Resultados: %d de %d respuestas correctas", respuestasCorrectas, respuestasTotales)
    conn.Write([]byte(resultados))

    // Enviar mensaje de finalización al cliente
    conn.Write([]byte("Fin Trivia"))

    // Esperar mensaje de finalización del cliente
    buffer := make([]byte, 1024)
    conn.Read(buffer)

    fmt.Println("Mensaje de finalización recibido del cliente. Cerrando conexión.")

    // Cerrar conexión TCP
    conn.Close()
}

