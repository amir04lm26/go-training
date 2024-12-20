package main

import "fmt"

func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)

	// Output: 0 1 2 3
	// * deferred functions runs in LIFO (Last In First Out) order
}
