package main

import "fmt"

// NOTE
/*
	Go does not have inheritance, but it allows composition through struct embedding.
	You can embed one struct inside another and inherit its methods and fields.
*/

type Person struct {
	name string
	age  uint8
}

type Employee struct {
	Person
	company string
}

func (person Person) greet() {
	fmt.Println("Hello, my name is", person.name)
}

func main() {
	employee := Employee{
		Person:  Person{name: "Amir", age: 26},
		company: "Snappshop",
	}

	// NOTE: This field is available as we're using composition in the Employee struct
	fmt.Println("Employee name:", employee.name)
	employee.greet()
}
