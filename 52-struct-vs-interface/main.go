package main

// NOTE: Struct vs Interface
/*
	In Go, struct and interface serve distinct purposes and are key components of the language's type system.
	Here's how they differ:
*/

// NOTE: 1. struct
/*
	Purpose: Defines a concrete data type that groups together fields (data).

	Usage: Used to represent real-world entities or data models with specific attributes.

	Definition:
	type Person struct {
		Name string
		Age  int
	}

	Key Points:
	Can have fields of any type, including other structs.
	Can have methods associated with it (using receivers).

	Requires explicit instantiation to use:
	p := Person{Name: "Alice", Age: 30}

	Useful for defining the structure of objects/data.
*/

// NOTE: 2. interface
/*
	Purpose: Defines a set of method signatures that a type must implement to be considered as satisfying the interface.

	Usage: Used to define behavior or contracts, enabling polymorphism.

	Definition:
	type Greeter interface {
		Greet() string
	}

	Key Points:
	Does not hold any data or implementation.
	Any type that implements all methods defined by the interface is said to satisfy it.
	Enables decoupling of code, allowing functions or components to operate on any type that satisfies the interface.

	func SayHello(g Greeter) {
		fmt.Println(g.Greet())
	}

	Can be implemented implicitly (no need for explicit declaration like implements in other languages):
	type Person struct {
		Name string
	}

	func (p Person) Greet() string {
		return "Hello, " + p.Name
	}

	Person now satisfies the Greeter interface.
*/

// NOTE: Comparison Table
/*
	Feature				struct											interface
	Purpose				Data representation								Behavioral abstraction
	Contains			Fields and methods								Method signatures
	Implementation		Explicitly defined								Implemented implicitly by any type
	Usage				Storing data and defining data structures		Defining contracts or polymorphic behavior
	Example				Person{Name: "Alice", Age: 30}					SayHello(g Greeter)
*/

//NOTE: When to Use
/*
	Use struct when you need to define and work with concrete data structures.
	Use interface when you want to abstract behavior and work with types polymorphically.

	Together, struct and interface allow Go to support a form of object-oriented programming
	that emphasizes composition over inheritance.
*/
