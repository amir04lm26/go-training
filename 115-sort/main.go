package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"
)

type person struct {
	name string
	age  int
}

// type people []person

// type sortBy string

// const (
// 	FieldName sortBy = "name"
// 	FieldAge  sortBy = "age"
// )

// func (people people) customSort(sortBy sortBy) {
// 	switch sortBy {
// 	case FieldName:
// 		fmt.Println("sorting by name")
// 	case FieldAge:
// 		fmt.Println("sorting by age")
// 	default:
// 		fmt.Println("invalid sort method")
// 	}
// }

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	// * custom sort
	p1 := person{"James", 32}
	p2 := person{"MoneyPenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	// people := people{p1, p2, p3, p4}
	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	// people.customSort(FieldName)
	// fmt.Println(people)
	// people.customSort(FieldAge)
	// fmt.Println(people)

	sort.Sort(ByAge(people)) // * ByAge(people) -> type conversion
	fmt.Println(people)

	slices.SortFunc(people, func(a, b person) int {
		if n := cmp.Compare(a.age, b.age); n != 0 {
			return n
		}
		return strings.Compare(a.name, b.name)
	})
	fmt.Println(people)
}
