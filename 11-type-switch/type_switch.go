package main

import (
	"fmt"
	"math"
)

func printType(variable interface{}) {
	// NOTE: switch is case insensitive
	switch value := variable.(type) {
	case string:
		println("Type is string:", value)
	case int:
		println("Type is int:", value)
	case uint16, uint32:
		fmt.Println("Type is either uint 16 or uint32:", value)
	default:
		println("Unknown type")
	}
}

func main() {
	var value uint8 = 23
	printType("Amir")
	printType(12)
	printType(math.Pi)
	printType(value)
	printType(uint16(123))
}
