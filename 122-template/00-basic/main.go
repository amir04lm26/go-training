package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// name := "Amir Zare"
	// * go run 00-basic/main.go "Amir Zare"
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)

	nf, err := os.Create("./00-basic/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
