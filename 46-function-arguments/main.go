package main

import "fmt"

// NOTE: Function argument passing mode
/*
	In Go, function arguments are not inherently immutable; rather, their mutability depends on
	the type of the argument being passed.

	Key Points:

	1.	Value Types:
		•	When you pass value types (like int, float64, bool, structs, etc.) to a function, they
			are passed by value. This means that a copy of the value is made, and modifications to
			this copy do not affect the original variable.

	2.	Reference Types:
		•	When you pass reference types (like slices, maps, channels, and pointers) to a function,
			the reference to the original data is passed. This means that modifications to the data
			can affect the original variable.

	3.	Using Pointers:
		•	If you want to modify a value type within a function, you can pass a pointer to that
			variable. This way, you can change the original value.
*/

// NOTE: Summary:
/*
	•	Value Types: Passed by value (immutable in terms of the original variable).
	•	Reference Types: Passed by reference (mutable, can modify the original data).
	•	Pointers: Allow modification of value types.

	So, whether function arguments are considered “mutable” or “immutable” in Go depends on their
	type and how they are passed to the function.
*/

func modifyValue(in int) {
	in = 10 // This does not affect the original variable
}

func valueType() {
	num := 5
	modifyValue(num)
	fmt.Println(num) // Output: 5
}

func modifySlice(slice []int) {
	slice[0] = 10 // This affects the original slice
}

func referenceType() {
	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println(mySlice) // Output: [10, 2, 3]
}

func modifyValuePtr(in *int) {
	*in = 10 // This affects the original variable
}

func pointerType() {
	num := 5
	modifyValuePtr(&num)
	fmt.Println(num) // Output: 10
}

func main() {
	valueType()
	referenceType()
	pointerType()
}
