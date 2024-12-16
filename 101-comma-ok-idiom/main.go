package main

import "fmt"

// NOTE: A Go map type looks like this:
/*
	map[KeyType]ValueType

	where KeyType may be any type that is comparable (more on this later), and ValueType may be
	any type at all, including another map!
*/

func main() {
	// var data map[string]int
	/*
		Map types are reference types, like pointers or slices, and so the value of m above is nil;
		it doesn’t point to an initialized map. A nil map behaves like an empty map when reading,
		but attempts to write to a nil map will cause a runtime panic; don’t do that. To initialize
		a map, use the built in make function:
	*/

	// data = make(map[string]int)
	/*
		The make function allocates and initializes a hash map data structure and returns a map
		value that points to it. The specifics of that data structure are an implementation detail
		of the runtime and are not specified by the language itself. In this article we will focus
		on the use of maps, not their implementation.
	*/

	// data := make(map[string]int)
	// data["age"] = 27

	data := map[string]int{
		"age": 27,
	}

	// NOTE: age will be the zero-value if the key doesn't exists
	if age, ok := data["age"]; ok {
		fmt.Println("age:", age)
	}
}
