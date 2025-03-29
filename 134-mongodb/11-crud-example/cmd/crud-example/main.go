package main

import (
	"crud-example/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
