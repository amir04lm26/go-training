package main

import "fmt"

/*
	In Go, both map and struct are used to store collections of data, but they differ in structure,
	use cases, performance, and flexibility. Here’s a detailed comparison:

	1. Definition

		•	map: A built-in data structure that represents a collection of key-value pairs.
		•	struct: A composite data type used to group fields into a meaningful entity with
			predefined names and types.

	2. Use Cases

		•	map:
		•		When keys are not known at compile time.
		•		Dynamic data access where you need to look up elements frequently by key.
		•		Examples: Lookup tables, JSON parsing, dictionaries.
		•	struct:
		•		When fields are fixed and known at compile time.
		•		Better for representing strongly-typed data models (e.g., a User or Product).
		•		Used to model entities with well-defined attributes.

	3. Syntax

		Map Example
		m := map[string]int{
			"apple":  5,
			"banana": 3,
		}
		fmt.Println(m["apple"]) // 5

		Struct Example
		type Fruit struct {
			Name  string
			Count int
		}
		f := Fruit{"apple", 5}
		fmt.Println(f.Name) // apple

	4. Performance Comparison

		•	Struct:
		•		Faster than a map because field access is statically resolved (no runtime lookup).
		•		Memory layout is contiguous, making it more cache-friendly.

		•	Map:
		•		Slower because it involves hash-based lookups and has more overhead.
		•		Good for dynamic keys but less efficient for fixed data.

	5. Flexibility

		•	Struct:
		•		Fields are statically defined and cannot change at runtime.
		•		Works well with the type system, enabling compile-time safety and IDE autocomplete.

		•	Map:
		•		Allows dynamic keys and values.
		•		Values can be added, removed, or modified at runtime, offering more flexibility.

	6. When to Use What?

		•	Use struct if:
		•		The structure and fields are known at compile time.
		•		You want better performance and static type safety.
		•		You are dealing with data models (e.g., User, Product).

		•	Use map if:
		•		Keys or the number of fields are not known beforehand.
		•		You need dynamic lookups (e.g., map[string]interface{} for unmarshaling JSON).

	7. Key Differences in Access

		•	Struct:
				fmt.Println(f.Name)  // Direct access by field name

		•	Map:
				fmt.Println(m["apple"])  // Access by key

		•	With structs, field access is compile-time checked, but map key access is not,
			which can lead to runtime errors if a key doesn’t exist.

	8. Memory Usage

		•	Struct:
		•		Memory is allocated for all fields.
		•		More predictable and optimized for access patterns.
		•	Map:
		•		Memory is allocated dynamically.
		•		Can grow and shrink based on the number of elements.

	9. Concurrency Considerations

		•	Struct:
		•		Safe for read operations but not thread-safe for writes without synchronization.
		•	Map:
		•		Not thread-safe for concurrent read and write operations. You need to use sync.Mutex or sync.Map for safe access in goroutines.

	10. Example Comparison

		Using Struct:
			type Person struct {
				Name string
				Age  int
			}

			p := Person{Name: "Alice", Age: 25}
			fmt.Println(p.Name)  // "Alice"

		Using Map:
			person := map[string]interface{}{
				"Name": "Alice",
				"Age":  25,
			}
			fmt.Println(person["Name"])  // "Alice"


	Conclusion

		•	Use struct when you have a well-defined structure with known fields.
		•	Use map when you need flexibility with keys or need to store varying fields dynamically.

	Both have their place in Go development. Structs offer better performance and type safety,
	while maps provide more flexibility for dynamic data.
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.Name) // "Alice"

	person2 := map[string]interface{}{
		"Name": "Alice",
		"Age":  25,
	}
	fmt.Println(person2["Name"]) // "Alice"
}
