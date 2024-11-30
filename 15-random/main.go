package main

import (
	"fmt"
	"go/go-examples/15-random/greeting"
	"log"
)

func main() {
	log.SetPrefix("Greeting: ")
	log.SetFlags(0)

	message, err := greeting.Hello("Amir")

	if err != nil {
		log.Panic(err) // NOTE: Serious problem
	}

	fmt.Println(message)
}
