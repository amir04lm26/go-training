package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	fmt.Println("Server is started at port 8080")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(connection)
	}
}

func handle(connection net.Conn) {
	defer connection.Close()

	request(connection)

	respond(connection)
}

func request(connection net.Conn) {
	iterator := 0
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if iterator == 0 {
			// * request line
			method := strings.Fields(line)[0]
			fmt.Println("***METHOD", method)
		}
		if line == "" {
			// * headers are done
			break
		}
		iterator++
	}
}

func respond(connection net.Conn) {
	tpl, err := template.New("response").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<title>Welcome to my http server!</title>
		</head>
		<body>
			<h1>Http Server</h1>
			<hr />
			<h2>I'm sending this response using a tcp server.</h2>
		</body>
		</html>
	`)
	if err != nil {
		log.Panicln(err)
	}

	var bodyBuf bytes.Buffer
	err = tpl.ExecuteTemplate(&bodyBuf, "response", nil)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", bodyBuf.Len())
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")

	_, err = bodyBuf.WriteTo(connection)
	if err != nil {
		log.Panicln(err)
	}
}
