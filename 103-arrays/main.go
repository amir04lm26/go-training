package main

import "fmt"

func main() {
	// * array literal
	arr1 := [3]int{42, 43, 44}
	fmt.Println("arr1", arr1)

	arr2 := [...]string{"Amir", "Faeze"}
	fmt.Println("arr2", arr2)

	var arr3 [2]int
	arr3[0] = 1
	arr3[1] = 2
	fmt.Println("arr3", arr3)

	fmt.Printf("arr1: %T\n", arr1)
	fmt.Printf("arr2: %T\n", arr2)
	fmt.Printf("arr3: %T\n", arr3)

	// NOTE: Number of type can be store in the array is part of the definition of it's type
	// arr3 = arr2 // ! error

	fmt.Println("arr2's len", len(arr2))

	// * new scope
	{
		arr1 := [3]string{"1", "2", "3"}
		fmt.Println("arr1", arr1)
	}
}
