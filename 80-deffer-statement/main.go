package main

import (
	"fmt"
	"os"
)

// NOTE: Defer Statement
/*
	The defer statement in Go is used to ensure that a function call is performed later,
	typically for cleanup purposes, like closing files or releasing resources.
	Deferred calls are executed in LIFO (Last In, First Out) order when the surrounding
	function returns.
*/

// NOTE: Example: Deferred Cleanup

// NOTE: Explanation:
/*
	•	defer for cleanup: file.Close() is deferred until the end of main, ensuring that
		the file is closed properly after writing to it.
	•	Resource management: Defer is commonly used for releasing resources (e.g., files,
		network connections) after use.
*/

func main() {
	file, err := os.Create("./80-deffer-statement/example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello, Go!")
	fmt.Println("File written successfully")
}
