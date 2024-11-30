package main

import (
	"fmt"
	"strings"
)

func printType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("Type is string:", v)
	case int:
		fmt.Println("Type is int:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	value := 3

	switch value {
	case 1, 3:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default case")
	}

	stringValue := "Example"
	normalizedValue := strings.ToLower(stringValue) // Normalize to lowercase

	switch normalizedValue {
	case "example":
		fmt.Println("Matched example case")
	case "test":
		fmt.Println("Matched test case")
	default:
		fmt.Println("No match found")
	}

	printType("hello")
	printType(42)
	printType(3.14)
}
