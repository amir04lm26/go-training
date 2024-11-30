package main

import (
	"fmt"
	"go/go-examples/51-read-only-in-go/abstraction"
	"go/go-examples/51-read-only-in-go/encapsulate-in-struct"
	syncOnce "go/go-examples/51-read-only-in-go/sync-once"
)

// NOTE: ReadOnly in go
/*
	In Go, variables cannot be explicitly declared as readonly like in some other languages
	(e.g., const in JavaScript or readonly in C#).
	However, you can achieve a readonly-like behavior by using one of the following approaches:
*/

// NOTE: 1. Use a const for Compile-Time Constants
/*
	If the value is known at compile time and does not change, you can use a const.
	Constants in Go are inherently immutable.
*/
func compileTimeConstant() {
	const pi = 3.14
	fmt.Println(pi)
	// Pi = 3.14159 // This will cause a compilation error
}

// NOTE: 2. Encapsulate the Value in a Struct with an Unexported Field
/*
	For runtime values, encapsulate the data in a struct and make the field unexported.
	Provide a getter method to access the value, and do not expose any setter methods.
*/

func encapsulateInStruct() {
	readonlyValue := encapsulate.NewReadOnly(42)
	fmt.Println(readonlyValue.GetValue())
	// readonlyValue.value = 10 // Compilation error: readonlyValue.value is unexported

	// NOTE
	/*
		If your main function and the ReadOnly struct are in the same package (such as package main),
		you can directly modify the unexported field value.
	*/
}

// NOTE: 3. Use Interfaces for Abstraction
/*
	You can define an interface that only exposes read operations and implement it in a type.
*/

func interfaceForAbstraction() {
	readonlyData := abstraction.NewReadOnlyData(100)
	fmt.Println(readonlyData.GetValue())
	// readonlyData.(dataStruct).value = 200 // This would be unsafe and is discouraged
}

// NOTE: 4. Use sync.Once for Initialization
/*
	If you want to set a value once and make it readonly afterward, you can use the sync.Once mechanism.
*/

func withSyncOnce() {
	readOnly := &syncOnce.ReadOnly{}
	readOnly.SetValue(12)
	fmt.Println(readOnly.GetValue()) // 12

	readOnly.SetValue(100)           // Ignored, value is already set
	fmt.Println(readOnly.GetValue()) // 12
}

// NOTE: Summary
/*
	Use const for compile-time readonly values.
	Use encapsulation with unexported fields and getter methods for runtime values.
	Use sync.Once for initializing a value only once in a thread-safe manner.
*/

func main() {
	compileTimeConstant()
	encapsulateInStruct()
	interfaceForAbstraction()
	withSyncOnce()
}
