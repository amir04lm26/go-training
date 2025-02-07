package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	fmt.Println("Server is running on port :8080")

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("New connection:", connection.LocalAddr())

		io.WriteString(connection, "\nHello from tcp server\n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v\n", "Well, I hope!")

		connection.Close()
	}
}
