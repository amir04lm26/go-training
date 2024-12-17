package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) // Output: 0, 1, 2, 3, 4, 5
	fmt.Println(b) // Output: 0, 1, 2, 3, 4, 5

	a[0] = 7
	fmt.Println(a) // Output: 7, 1, 2, 3, 4, 5
	fmt.Println(b) // Output: 7, 1, 2, 3, 4, 5

	// NOTE: copy
	/*
	   The copy built-in function copies elements from a source slice into a destination slice.
	   (As a special case, it also will copy bytes from a string to a slice of bytes.)
	   The source and destination may overlap. Copy returns the number of elements copied,
	   which will be the minimum of len(src) and len(dst).
	*/
	c := make([]int, len(a))
	copy(c, a)
	fmt.Println(a) // Output: 7, 1, 2, 3, 4, 5
	fmt.Println(b) // Output: 7, 1, 2, 3, 4, 5
	fmt.Println(c) // Output: 7, 1, 2, 3, 4, 5

	a[0] = 2
	fmt.Println(a) // Output: 2, 1, 2, 3, 4, 5
	fmt.Println(b) // Output: 2, 1, 2, 3, 4, 5
	fmt.Println(c) // Output: 7, 1, 2, 3, 4, 5
}
