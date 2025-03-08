package main

import (
	"log"
	"mongo-basic/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
