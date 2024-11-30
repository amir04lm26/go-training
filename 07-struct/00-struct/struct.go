package main

import "fmt"

type Person struct {
	name string
	age  int
}

func birthday(person *Person) {
	(*person).age++
}

func main() {
	var person1 Person = Person{name: "Amir", age: 26}

	println("name:", person1.name)
	println("age:", person1.age)

	person1.age = 22
	fmt.Printf("Person one is %d years old\r\n", person1.age)

	birthday(&person1)
	fmt.Printf("Person one is %d years old\r\n", person1.age)
}
