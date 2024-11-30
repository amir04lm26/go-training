package main

import (
	"fmt"
	"reflect"
)

// NOTE: Reflection in Go
/*
	Reflection is a powerful feature in Go that allows you to inspect and manipulate types at runtime.
	The reflect package provides the necessary tools for working with reflection.
*/

// NOTE: Basic Reflection
/*
	The two primary types in the reflect package are:
		•	reflect.Type: Describes the type of a value.
		•	reflect.Value: Describes the actual value.
*/

// NOTE: Explanation:
/*
	•	reflect.TypeOf returns the type of the value (float64 in this case).
	•	reflect.ValueOf returns a reflect.Value which holds the actual value.
*/

func basicReflection() {
	var num float64 = 3.14

	fmt.Println("type:", reflect.TypeOf(num))
	fmt.Println("value:", reflect.ValueOf(num))

	value := reflect.ValueOf(num)
	fmt.Println("kind is float:", value.Kind() == reflect.Float64)
	fmt.Println("value is:", value.Float())
}

// NOTE: Modifying Values with Reflection
/*
	To modify a value using reflection, the value must be passed as a pointer.
*/

func modifyValueWithReflection() {
	var num float64 = 3.4
	pointer := reflect.ValueOf(&num)
	value := pointer.Elem()

	fmt.Println("Pointer:", pointer)
	fmt.Println("Old value:", value.Float())
	value.SetFloat(7.1)
	fmt.Println("New value:", num)
}

// NOTE: Explanation
/*
	•	Pointer: You need to pass a pointer to the value to modify it (reflect.ValueOf(&num)).
	•	Elem(): This dereferences the pointer to get the actual value.
	•	SetFloat(): This modifies the underlying value.
*/

/*
	This concludes the section on Reflection in Go. It’s a powerful feature but should be used
	sparingly, as it can make the code harder to understand and slower.
*/

func main() {
	basicReflection()
	modifyValueWithReflection()
}
