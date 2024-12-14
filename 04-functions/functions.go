package main

import "fmt"

// NOTE: When we define function, it gets parameters
// NOTE: When we use the function, it gets arguments

func greet(name string) {
	println("Hello, ", name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}

func main() {
	greet("Amir")

	added := add(23, 62)
	println("Some up value:", added)

	divided, err := divide(10, 2)
	if err != nil {
		println("Error:", err)
	} else {
		println("divided:", divided)
	}

	divideRes, divideErr := divide(10, 0)
	if divideErr != nil {
		fmt.Println("Error:", divideErr)
	} else {
		println("divided:", divideRes)
	}

	func() {
		value := "IIFE"
		fmt.Printf("I believe this is an immediately invoked function expression (%v)\r\n", value)
	}()
	// value = "error" > error > undefined
}
