package main

import "fmt"

// NOTE: method sets
/*
	The idea of the method set is integral to how interfaces are implemented and used in Go.
		- The method set of a type T consists of all methods with receiver type T.
		  These methods can be called using variables of type T.

		- The method set of a type *T consists of all methods with receiver *T or T
		  These methods can be called using variables of type *T.
		  it can call methods of the corresponding non-pointer type as well
*/

type dog struct {
	name string
}

func (dog dog) walk() {
	fmt.Println("My name is", dog.name, "and I'm walking")
}

func (dog *dog) run() {
	dog.name = "Rover"
	fmt.Println("My name is", dog.name, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(youngin youngin) {
	youngin.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) // * error

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)
}
