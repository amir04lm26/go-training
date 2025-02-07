package main

import (
	"bufio"
	"fmt"
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

		go handle(connection)
	}
}

func handle(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "I head you said %s\n", line)
	}

	// * We never get here
	// * we have an open stream connection
	// * How does the above reader know when it's done?
	fmt.Println("Code got here")
}
