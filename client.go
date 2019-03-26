package main

import (
	"bufio"
	"fmt"
	"github.com/ArielPierot/tmqp-api-client"
	"os"
)

func main() {

	conn := tmqp.Handshake()
	fmt.Println("A conexão foi estabelecida, você já pode digitar os comandos.")

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=> ")
		queue, _ := r.ReadString('\n')
		tmqp.NewQueue(conn, queue)

		tmqp.Consume(conn, queue)

		for {
			fmt.Print("=> ")
			message, _ := r.ReadString('\n')
			tmqp.Publish(conn, "Ariel", queue, message)
		}

	}

}
