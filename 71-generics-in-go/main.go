package main

import "fmt"

// NOTE: Generics in Go (since Go 1.18)
/*
	Generics allow you to write code that works with any data type, increasing code reusability
	and type safety.
*/

// NOTE: Example: Generic Function

// NOTE: Explanation:
/*
	•	[T int | float64]: Declares that T can be of type int or float64.
	•	Min(): This function works for both integers and floating-point numbers.
*/

// * Generic function to find the minimum value in a slice
func Min[T int | float64](num1, num2 T) T {
	if num1 < num2 {
		return num1
	}
	return num2
}

func main() {
	fmt.Println(Min(3, 5))     // * Works with int
	fmt.Println(Min(3.2, 5.1)) // * Works with float64
}
