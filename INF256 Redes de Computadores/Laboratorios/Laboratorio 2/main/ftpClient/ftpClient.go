package main

import (
	"os"

	"github.com/secsy/goftp"
)

const (
	ftpServerAddr = "192.168.1.188" //actualizar si es necesario
	ftpServerPath = "." //mantener esta línea para trabajar en la carpeta raíz del servidor
)

func main() {
	config := goftp.Config{
		User:     "10", //utilice el usuario asignado a su grupo
		Password: ".", //utilice la contraseña obtenida durante la interacción con el servidor TCP
		// Como no se logro obtener la contraseña, se dejará un punto
	}
	client, err := goftp.DialConfig(config, ftpServerAddr)
	if err != nil {
		panic(err)
	}

	test, err := os.Create("Preguntas.txt") //utilice el nombre del archivo proporcionado en el laboratorio
	if err != nil {
		panic(err)

	}

	err = client.Retrieve("./Preguntas.txt", test) //utilice el nombre del archivo proporcionado en el laboratorio

	if err != nil {
		panic(err)
	}

	bigFile, err := os.Open("tcp.txt")
		if err != nil {
		panic(err)
	}

	err = client.Store("./tcp.txt", bigFile)
		if err != nil {
		panic(err)
	}
}
