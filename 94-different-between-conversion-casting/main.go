package main

import (
	"fmt"
	"strconv"
)

// NOTE: Difference Between Conversion and Casting

// NOTE: 1. Casting:
/*
	Definition: Casting typically refers to directly reinterpreting a value's type without changing
	its actual data representation in memory.
	Behavior: Casting is more common in languages like C and C++ and can be unsafe, as it might
	lead to unexpected behavior if the types are incompatible (e.g., casting a float to a pointer).

	Example in C:
	```main.c
	float f = 3.14;
	int i = (int)f; // Casts the float value to an integer (truncates).
	```
*/

// NOTE: Conversion:
/*
	Definition: Conversion explicitly transforms one type into another in a safer manner, ensuring
	compatibility between the two types.
	Behavior: Conversion may involve some transformation or reallocation of memory, as the data
	format might need to change.

	Example in Go:
	```go
	var i int = 42
	var f float64 = float64(i) // Converts int to float64
	```
*/

// NOTE: Conversion in Go
/*
	In Go, conversion is used to explicitly transform one type to another.
	Unlike casting in other languages, Go does not allow implicit type coercion,
	so you must explicitly use conversion whenever you want to change types.

	Key Characteristics of Conversion in Go:
	1. Explicitness: You must use a conversion function to change a type.
	2. Safety: Go ensures that conversions are valid at compile time.
	3. Syntax:
		The target type is written as a function.
		Example: float64(variable) converts a variable to float64.
*/

// NOTE: Examples of Conversion in Go:

// NOTE: In Summary:
/*
	Casting: Reinterprets data types without altering representation; often unsafe.
	Conversion: Explicitly transforms data types safely and is strongly enforced in Go.
	In Go: Conversion is a deliberate and explicit transformation of types that is safe
	and type-checked at compile time.

	In Go, conversions are always explicit, and the language enforces strong type safety
	at compile time.

	Go does not have type coercion. Go is a statically typed language,
	meaning you must explicitly convert values from one type to another when needed.
	This strict typing helps prevent subtle bugs and makes the language behavior more predictable.
*/

// NOTE: Type Assertions (Not Coercion)
/*
	Go does have a concept of type assertions, which are used to work with interface types,
	but this is not type coercion. Type assertions are for runtime type checking.
*/

// NOTE: Example of Type Assertion:
/*
	package main

	import "fmt"

	func main() {
		var i interface{} = "hello"

		// Type assertion to convert `i` to a string
		s, ok := i.(string)
		if ok {
			fmt.Println(s) // Output: hello
		}
	}
*/

// NOTE: Summary
/*
	Go does not have implicit type coercion.
	Explicit type conversion is required for any type changes.
	This strictness ensures type safety, predictability, and clarity in code.
*/

func main() {
	// NOTE: 1. Numeric Types:
	var i int = 42
	var f float64 = float64(i) // * int to float64
	var u uint = uint(f)       // * float64 to uint

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", u, u)

	// NOTE: 2. String and Byte Slice:
	var s string = "hello"
	var b []byte = []byte(s)   // * string to byte slice
	var str string = string(b) // * byte slice back to string

	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", str, str)

	// NOTE: 3. Custom Types: You can also convert between types if they have the same underlying type.
	type Celsius float64
	type Fahrenheit float64

	var tempC Celsius = 25
	var tempF Fahrenheit = Fahrenheit(tempC*9.5 + 32) // * Celsius to Fahrenheit

	fmt.Printf("%v %T\n", tempC, tempC)
	fmt.Printf("%v %T\n", tempF, tempF)

	// NOTE: When Conversion is Not Possible:
	/*
		Go prevents conversions that could lead to unsafe or invalid states, such as:
	*/
	var x int = 64
	// * This does not convert an int to a string (results in invalid ASCII character)
	var y string = string(x)

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// NOTE: Instead, you might need a library function, like strconv:
	var z string = strconv.Itoa(x) // * Proper int-to-string conversion
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", z, z)
}
