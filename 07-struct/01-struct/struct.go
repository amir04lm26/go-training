package main

import "fmt"

type Address struct {
	city  string
	state string
}

// Struct with Embedded Structs
type Employee struct {
	name    string
	age     uint8
	address Address
}

func main() {
	// NOTE: Can't be a const
	employee1 := Employee{
		name: "Amir",
		age:  26,
		address: Address{
			city:  "Tehran",
			state: "Tehran",
		},
	}

	fmt.Printf("%s lives in %s.\r\n", employee1.name, employee1.address.city)
	fmt.Printf("%s is %d.\r\n", employee1.name, employee1.age)
}
