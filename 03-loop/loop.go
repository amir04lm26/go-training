package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: classic for loop
	println("# classic for loop")
	for iterator := 0; iterator < 5; iterator++ {
		println("iterator:", iterator)
	}

	// NOTE: While like loop
	println("\r\n# While like loop")
	iterator := 5
	for 0 < iterator {
		println("iterator:", iterator)
		iterator--
	}

	// NOTE: continue -> takes to next iteration
	for iterator := 0; iterator < 20; iterator++ {
		if iterator%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers:", iterator)
	}

	// NOTE: Infinite loop
	println("\r\n# Infinite loop")
	for {
		println("This is an infinite loop.")
		time.Sleep(1 * time.Second)
	}
}
