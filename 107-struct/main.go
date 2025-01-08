package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// * embedding a struct in another struct
type employee struct {
	// person person
	person
	id    int
	first string
}

func main() {
	// * struct/composite literal
	// p1 := person{"Amir", "Zare", 27}
	p1 := person{first: "Amir", last: "Zare", age: 27}
	fmt.Printf("%v,\t %T,\t %#v\n", p1, p1, p1)

	e1 := employee{p1, 123, "first"}
	fmt.Printf("%#v\n", e1)

	// NOTE: Inner type promotion: the property will be accessAble from the upper level
	// NOTE: Also for methods
	e2 := employee{person{"Amir", "Zare", 27}, 123, "new"}
	fmt.Println(e2, e2.first, e2.person.first)
	fmt.Println(e2, e2.age)

	// * anonymous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{"Amir", "Zare", 27}
	fmt.Printf("%v,\t %T,\t %#v\n", p2, p2, p2)

	// NOTE: You can not assign a named type to another named type
	// NOTE: For better performance lay the fields out from largest to smallest, e.g int64, float32, bool
}
