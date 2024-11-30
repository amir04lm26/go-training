package main

import "fmt"

func main() {
	const age = 26

	if age >= 18 {
		fmt.Println("Eligible to vote.")
	} else {
		fmt.Println("Not eligible to vote.")
	}
}
