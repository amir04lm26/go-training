package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()

	byteStr, err := io.ReadAll(connection)
	if err != nil {
		defer connection.Close()
	}

	fmt.Println(string(byteStr))
}
