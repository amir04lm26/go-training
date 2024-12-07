package main

import (
	"errors"
	"fmt"
)

// NOTE: Error Wrapping with fmt.Errorf and errors.Is
/*
	Starting from Go 1.13, Go introduced error wrapping, allowing you to attach context to errors.
*/

// NOTE: Example: Error Wrapping

// NOTE: Explanation:
/*
	•	%w: In fmt.Errorf, %w wraps an error so it can be checked later.
	•	errors.Is(): Compares the error to a known error like ErrNotFound.
*/

var ErrNotFound = errors.New("not found")

func findValue(id int) error {
	return fmt.Errorf("failed to find value: %w", ErrNotFound)
}

func main() {
	err := findValue(42)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("The error is 'not found'")
	}
}
