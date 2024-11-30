package main

import "fmt"

type Greeter interface {
	greet()
}

type Person struct {
	name string
	age  uint8
}

type Robot struct {
	id string
}

func (person Person) greet() {
	fmt.Println("Hello, my name is", person.name)
}

func (robot Robot) greet() {
	fmt.Println("Hello, I'm robot", robot.id)
}

// NOTE: Custom-Type
/*
	In Go, you cannot directly implement an interface on a basic type such as float64.
	However, you can achieve similar behavior by creating a custom type based on the float64 type
	and then implementing the interface methods on that custom type.
*/
type FloatInterface interface {
	Double() float64
}

type CustomFloat float64

func (num CustomFloat) Double() float64 {
	return float64(num * 2)
}

func customType() {
	var num CustomFloat = 5.5
	var myInterface FloatInterface = num

	fmt.Println("Double:", myInterface.Double())
	fmt.Println("num:", num)
}

func main() {
	var greeter Greeter

	greeter = Person{name: "Amir", age: 26}
	greeter.greet()

	greeter = Robot{id: "R2-02"}
	greeter.greet()

	customType()
}
