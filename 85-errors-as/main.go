package main

import (
	"errors"
	"fmt"
	"os"
)

// NOTE: Wrap Errors for Context
/*
	When returning an error, wrap it with additional context using the %w verb in fmt.Errorf.
	This allows you to unwrap the error later if needed.
*/

// NOTE: Example:
func readFile(filePath string) error {
	_, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("readFile: %w", err)
	}
	return nil
}

func main() {
	if err := readFile("non-existent.txt"); err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Println("Path error:", pathErr)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
