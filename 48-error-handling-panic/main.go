package main

import "fmt"

// NOTE: Panic and Recover
/*
	In Go, panic is a way to stop the normal execution of a program. It is used in rare situations
	like critical errors, where recovery is either impossible or undesirable.
*/

// NOTE: Panic
/*
	When a program encounters a panic, it prints the error message and stack trace and then exits.
*/

// NOTE: Fatal
/*
	When we use Fatal the program exists using os.exit, but in panic we can use recover
*/

// NOTE: Explanation
/*
	•	The panic function is called when an invalid operation occurs (in this case, dividing by zero).
	•	Once panic is triggered, the program immediately starts to unwind the call stack.
*/

func divide(num1, num2 float64) float64 {
	if num2 == 0 {
		panic("cannot divide by zero") // NOTE: Trigger a panic
	}
	return num1 / num2
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0)) // NOTE: This will cause the program to panic
	fmt.Println("This will not be printed")
}
