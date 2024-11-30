package main

type Person struct {
	name string
	age  uint8
}

func (person Person) greet() {
	println("Hello, my name is", person.name, "\r\nI'm", person.age, "years old.")
}

func (person *Person) birthday() {
	(*person).age++
}

func main() {
	person1 := Person{name: "Amir", age: 26}
	person1.greet()

	person1.birthday()
	person1.greet()
}
