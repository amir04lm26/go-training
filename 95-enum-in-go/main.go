package main

import "fmt"

// NOTE: ENUM in Go
/*
	Go doesn't have a built-in enum keyword like some other programming languages
	(e.g., Java or C#), but it provides a way to achieve similar functionality using
	constant blocks and the iota identifier. This is a common pattern in Go for creating
	enumerated values.
*/

// NOTE: How to Create Enums in Go
/*
	In Go, you typically define an "enum-like" construct using a const block.
	Here's how it works:
*/

// NOTE: In this example:
/*
	iota is a special identifier that generates a sequence of incrementing integer values,
	starting from 0 within a const block.
*/

const (
	Red   = iota // * iota starts at 0 and increments automatically
	Green        // * 1
	Blue         // * 2
)

func basicEnumExample() {
	fmt.Println(Red, Green, Blue) // * Output: 0 1 2
}

// NOTE: Custom Enum Values
/*
	You can assign specific values or manipulate iota to create custom enumerations.
*/

// NOTE: Example with Custom Values

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// NOTE: Example with Custom Starting Value
const (
	ErrorLevel   = iota + 1 // Start from 1
	WarningLevel            // 2
	InfoLevel               // 3
	DebugLevel              // 4
)

// NOTE: Enum with Specific Types
/*
	To create a strongly-typed enum, you can define a type and associate the constants with it.
	This improves type safety.
*/

// NOTE: Example with Typed Enum

type Color int // * Define a custom type

const (
	White Color = iota // * Assign constants to the custom type
	Gray
	Black
)

func withTypedEnumExample() {
	var color Color = Red
	fmt.Println(color) // Output: 0

	// NOTE: Compare enums
	if color == Red {
		fmt.Println("Color is Red") // * Output: Color is Red
	}
}

// NOTE: Using Enums with String Representations
/*
	Go doesn't automatically convert enum values to strings, but you can implement this
	manually using a map or a switch statement.
*/

// NOTE: Example with String Representation

type CustomColor int

const (
	Yellow CustomColor = iota
	Orange
	Pink
)

func (color CustomColor) String() string {
	switch color {
	case Yellow:
		return "Yellow"
	case Orange:
		return "Orange"
	case Pink:
		return "Pink"
	default:
		return "Unknown"
	}
}

func enumWithStringRepresentation() {
	fmt.Println(Yellow) // Output: Yellow
	fmt.Println(Orange) // Output: Yellow
	fmt.Println(Pink)   // Output: Yellow
}

// NOTE: When to Use Enums in Go
/*
	Enums are great for fixed sets of related constants (e.g., days of the week, states, levels, etc.).
	Using a custom type for enums improves type safety and makes your code more self-documenting.
*/

// NOTE: Summary
/*
	Go doesn't have a built-in enum keyword.
	You can emulate enums using const and iota.
	Strongly-typed enums can be created by defining a custom type.
	You can add string representations manually for better usability.
*/

func main() {
	basicEnumExample()
	withTypedEnumExample()
	enumWithStringRepresentation()
}
