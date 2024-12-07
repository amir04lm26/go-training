package main

import "fmt"

// NOTE: Custom Error Types
/*
	In Go, you can define custom error types by implementing the Error method on a type,
	which allows you to create structured and meaningful error messages.
*/

// NOTE: Example: Custom Error Type

// NOTE: Explanation:
/*
	•	Custom error struct: MyError is a struct that holds details about the error.
	•	Implementing Error(): The Error method formats and returns a string representation
		of the error.
*/

type MyError struct {
	When string
	What string
}

func (error *MyError) Error() string {
	return fmt.Sprintf("Error occurred at %s: %s", error.When, error.What)
}

func riskyOperation() error {
	return &MyError{
		When: "2024-09-13",
		What: "Something went wrong",
	}
}

func main() {
	err := riskyOperation()
	if err != nil {
		fmt.Println(err)
	}
}
