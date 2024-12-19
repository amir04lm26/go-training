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
			values types: int, float64, bool, string, uint, structs, arrays and Custom Defined
			Value Types

	2.	Reference Types:
		•	When you pass reference types (like slices, maps, channels, and pointers) to a function,
			the reference to the original data is passed. This means that modifications to the data
			can affect the original variable.
			reference types: slices, maps, channels, pointers, interfaces and functions

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

// NOTE: Value Types vs Reference Types
/*
	Aspect				Value Types							Reference Types
	Behavior			Data is copied during assignment.	Reference to the same data is passed.
	Examples			int, float64, struct, array			slice, map, pointer, channel
	Memory Allocation	Stored directly in memory (stack).	Stored in memory indirectly (heap).
	Modification		Does not affect original variable.	Affects original variable.
*/

// NOTE: Practical Considerations
/*
	1. Use Pointers to Avoid Copying Large Data:
		For large structs or arrays, copying can be expensive. Instead, use pointers for efficiency.
		Example:
		```go
		type Person struct {
			Name string
			Age  int
		}

		func updateAge(p *Person) {
			p.Age = 40
		}

		func main() {
			p1 := Person{"Alice", 30}
			updateAge(&p1)
			fmt.Println(p1.Age) // Output: 40
		}
		```

	2. Slices and Maps Are Reference Types:
		Even though slices and maps are "based" on arrays, they behave like reference types in Go
		because they internally manage a reference to the underlying data.
*/

// NOTE: Summary
/*
	Value types in Go include basic types (int, float64, bool, string), arrays, and structs.
	When assigned or passed to functions, they are copied.
	Use pointers to work with value types efficiently when modifying large or complex data structures.
*/

// NOTE: Primitive types

// NOTE: Characteristics of Primitive Types:
/*
	1. Basic and Built-in:
		Primitive types are predefined by the programming language.
		Example: int, float, bool, string.

	2. Single Value Storage:
		Each variable of a primitive type stores a single piece of data (e.g., a number, character,
		or boolean value).

	3. Immutable in Some Cases:
		In many languages, primitive types are immutable, meaning their value cannot be
		changed once assigned (e.g., string in Go).

	4. Efficient:
		Primitive types are highly optimized for performance and often map directly to
		machine-level instructions.

	5. Independent:
		They don't depend on other types or structures.
*/

/*
	Go has a set of primitive types:

	Boolean: bool (true or false)
	Numeric:
	Integer: int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64
	Floating-point: float32, float64
	Complex numbers: complex64, complex128
	String: string
	Byte: byte (alias for uint8)
	Rune: rune (alias for int32, used for Unicode characters)
*/

// NOTE: Primitive Types vs. Non-Primitive Types
/*
	Aspect			Primitive Types				Non-Primitive Types
	Complexity		Basic and simple			Can be made up of multiple values
	Examples		int, float, bool, string	array, slice, map, struct, class
	Storage			Holds direct data			Holds references to data structures
	Performance		Fast and memory-efficient	Slower due to complexity
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
