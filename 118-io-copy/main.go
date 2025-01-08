package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("./118-io-copy/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := strings.NewReader("Amir Zare")

	io.Copy(file, reader)
}
