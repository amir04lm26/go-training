package main

import (
	"fmt"
)

func main() {
	var name string = "Amir"
	var age int8 = 26
	fmt.Println("Name:", name, ", Age:", age)

	const nName, nAge, _ = "Amir", 27, 3
	// * '_' -> Blank Identifier
	fmt.Println("Name:", nName, ", Age:", nAge)

	fmt.Printf("27 as binary, %b\n", nAge)
	fmt.Printf("27 as hexadecimal, %x\n", nAge)
	fmt.Printf("27 as hexadecimal, %#x\n", nAge)

	// ASCII - 1 byte
	// UNICODE - byte
	// utf-8 - up to 4 byte
	fmt.Println("Emoji: 😉")

	var g int // * All types have a zero-value
	var d string
	var e bool
	var f float64
	fmt.Println("g:", g, "d:", d, "e:", e, "f:", f)

	// NOTE: Zero-Values
	/*
		boolean 		false
		int     		0
		float     		0.0
		string   		""
		pointers 		nil
		functions 		nil
		interfaces		nil
		slices 			nil
		channels 		nil
		map 			nil
	*/

	// NOTE: General Guideline: User short decoration operator for all of variables if possible
	// NOTE: We cam use it only inside a function
	// NOTE: There is no variable hoisting in go
	// NOTE: When we export a variable it should start with a Capital letter
	/*
		In Go, a name is exported if it begins with a Capital letter. For example, Pizza is a
		exported name, as is Pi, which is exported from the math package.
	*/

	// NOTE: Short variable declaration
	/*
		Outside a function, every statement begins with a keyword (var, func, and so on) and
		so the := construct is not available.
	*/
	lastName := "Zare"
	fmt.Println("Last Name:", lastName)

	const wife = "Faeze"
	fmt.Println("Wife:", wife)

	const pi float64 = 3.14159
	const condition bool = true
	fmt.Println("pi:", pi, ", condition:", condition)
}
