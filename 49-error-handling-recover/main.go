package main

import "fmt"

// NOTE: Recover
/*
	The recover function allows you to regain control of a panicking goroutine and prevent the
	program from crashing.
*/

// NOTE: Explanation
/*
	•	Defer: The defer keyword schedules a function to be called after the surrounding function returns.
	•	Recover: The recover function is used to catch the panic and prevent the program from crashing. It allows for graceful handling
		of errors caused by panic.
*/

func divide(num1, num2 float64) float64 {
	if num2 == 0 {
		panic("cannot divide by zero")
	}
	return num1 / num2
}

func safeDivide(num1, num2 float64) (result float64) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recovering from panic:", rec)
			result = 0 // Set default value on panic
		}
	}()
	return divide(num1, num2)
}

func main() {
	fmt.Println("Result:", safeDivide(10, 2))
	fmt.Println("Result:", safeDivide(10, 0))
}
