package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	name := "Go"
	version := 1.20
	year := 2024
	usage := 1_000_000

	println("Simple log")
	println("Language:", name, "Version:", version)

	fmt.Println("Language:", name, "Version:", version)

	fmt.Printf("Language: %s, Version: %.2f\n", name, version)
	fmt.Printf("Language: %s, Version: %.3f\n", name, version)
	fmt.Printf("Year: %d\n", year)
	fmt.Printf("Usage: %d\n", usage)
	fmt.Printf("Language: %s, Version: %.2f, Year: %d, Usage: %d\n", name, version, year, usage)

	// NOTE
	/*
		%f: Always uses fixed-point notation (e.g., 123.456).
		%e: Always uses scientific notation (e.g., 1.23456e+02).
		%g: Dynamically chooses between %f and %e depending on the value and length of the number. It aims to give the most compact representation.
	*/
	value1 := 123456.789
	value2 := 0.000012345
	value3 := 123.456

	fmt.Printf("%g\n", value1)   // Will use fixed-point notation
	fmt.Printf("%g\n", value2)   // Will use scientific notation
	fmt.Printf("%g\n", value3)   // Will use fixed-point notation
	fmt.Printf("%.2g\n", value1) // Limits to 2 significant digits

	// NOTE: Using Sprintf to format a string
	personName := "Amir"
	personAge := 26
	printValue := fmt.Sprintf("My name is %s and I am %d years old.", personName, personAge)

	fmt.Println(printValue)

	// NOTE
	/*
		•	%v: Default value formatting for most types.
		•	%+v: For structs, it includes field names along with their values.
		•	%#v: Prints the Go syntax of the value, useful for debugging and seeing exact struct representation.
		•	%T: Prints the type of the variable (not the value, just the type).
	*/
	myInt := 42
	myString := "hello"
	myFloat := 3.14
	myBool := true

	fmt.Printf("int: %v\n", myInt)
	fmt.Printf("string: %v\n", myString)
	fmt.Printf("float: %v\n", myFloat)
	fmt.Printf("bool: %v\n", myBool)

	fmt.Printf("myInt: %v, type: %T\n", myInt, myInt)
	fmt.Printf("myString: %v, type: %T\n", myString, myString)
	fmt.Printf("myFloat: %v, type: %T\n", myFloat, myFloat)
	fmt.Printf("myBool: %v, type: %T\n", myBool, myBool)

	person := Person{Name: "Amir", Age: 26}
	fmt.Printf("Person: %v\n", person)
	// NOTE: %+v: This variant includes the field names when printing structs.
	fmt.Printf("Person: %+v\n", person)
	// NOTE: %#v: This variant prints the value in Go syntax, useful for debugging.
	fmt.Printf("Person: %#v\n", person)

	myArr := []int{1, 2, 3} // NOTE: Slice
	myMap := map[string]int{"a": 1, "b": 2}
	fmt.Printf("Slice: %v\n", myArr)
	fmt.Printf("Map: %v\n", myMap)

	myPointer := &myInt
	fmt.Printf("Pointer: %v\n", myPointer)

	complexNum := complex(2, 3)
	fmt.Printf("Complex: %v\n", complexNum)

	err := errors.New("something went wrong")
	fmt.Printf("Error: %v\n", err)
}
