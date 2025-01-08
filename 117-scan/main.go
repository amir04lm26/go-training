package main

import (
	"fmt"
	"log"
)

func main() {
	var name string

	fmt.Print("Please enter your name: ")
	// _, err := fmt.Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if _, err := fmt.Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", name)
}
