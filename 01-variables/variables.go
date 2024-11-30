package main

import "fmt"

func main() {
	var name string = "Amir"
	var age int8 = 26
	fmt.Println("Name:", name, ", Age:", age)

	lastName := "Zare"
	fmt.Println("Last Name:", lastName)

	const wife = "Faeze"
	fmt.Println("Wife:", wife)

	const pi float64 = 3.14159
	const condition bool = true
	fmt.Println("pi:", pi, ", condition:", condition)
}
