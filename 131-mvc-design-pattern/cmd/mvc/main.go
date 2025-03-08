package main

import (
	"log"
	"mvc-design-pattern/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
