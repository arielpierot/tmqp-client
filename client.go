package main

import (
	"fmt"
	"github.com/ArielPierot/tmqp-api-client"
)

func main() {

	var comando int8

	conn := tmqp.Handshake()

	for {
		fmt.Print("Informe o comando desejado: ")
		fmt.Scanf("%d\n", &comando)
		tmqp.SendPackage(conn, comando)
	}

}
