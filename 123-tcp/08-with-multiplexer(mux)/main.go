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

	method, uri := request(connection)

	// NOTE: Mux should resolve the desire response handler fn instead of this
	respond(connection, method, uri)
}

func request(connection net.Conn) (method, uri string) {
	iterator := 0
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if iterator == 0 {
			// * request line
			fields := strings.Fields(line)
			method = fields[0]
			uri = fields[1]
			fmt.Println("***METHOD", method)
			fmt.Println("***URI", uri)
		}
		if line == "" {
			// * headers are done
			break
		}
		iterator++
	}

	return
}

func respond(connection net.Conn, method, uri string) {
	isHomePage := method == "GET" && uri == "/"
	tpl, err := template.New("templates").Parse(`
		{{define "not-found"}}
			<h2>The page you're looking for is not found!</h2>
		{{end}}

		{{define "home"}}
			<h2>I'm sending this response using a tcp server.</h2>
		{{end}}
		
		{{define "response"}}
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<title>Welcome to my http server!</title>
			</head>
			<body>
				<h1>Http Server</h1>
				<hr />
				<h2>Request: {{.URI}}</h2>
				{{if .IsHomePage}}
					{{template "home"}}
				{{else}}
					{{template "not-found"}}
				{{end}}
			</body>
			</html>
		{{end}}
	`)
	if err != nil {
		log.Println(err)
	}

	var bodyBuf bytes.Buffer
	err = tpl.ExecuteTemplate(&bodyBuf, "response", struct {
		IsHomePage bool
		URI        string
	}{isHomePage, uri})
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", bodyBuf.Len())
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")

	_, err = bodyBuf.WriteTo(connection)
	if err != nil {
		log.Println(err)
	}
}
