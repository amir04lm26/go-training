package main

import (
	"encoding/json"
	"fmt"
)

// NOTE: Struct Tags in Go
/*
	Struct tags provide metadata for struct fields. These are typically used by libraries
	(like encoding or database libraries) to customize how struct fields are processed.
*/

// NOTE: Example: Using Struct Tags

// NOTE: Explanation:
/*
	•	Struct tag: The field Name has a struct tag json:"name", which tells the JSON package to
		use "name" as the key in the JSON object.
	•	Customization: Struct tags are used for customizing the behavior of encoding, database
		field mapping, and other operations.
*/

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Amir", Age: 27}
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data:", string(data))
}
