package main

import "fmt"

func main() {
	// * compose/slice literal
	xs := []string{"Amir", "Faeze"}
	fmt.Println("xs", xs)

	xs = append(xs, "married", "1403")
	fmt.Println("xs", xs)

	for iterator := 0; iterator < len(xs); iterator++ {
		fmt.Println(xs[iterator])
	}

	for key, value := range xs {
		fmt.Printf("%v: %v\n", key, value)
	}

	// NOTE: slicing a alice
	// * [inclusive:exclusive)
	fmt.Println(xs[0:2]) // * up to 2
	fmt.Println(xs[:2])  // * up to 2
	fmt.Println(xs[2:])  // * from 2
	fmt.Println(xs[:])   // * everything

	fmt.Println(xs[:len(xs)-1]) // * not the last one
	fmt.Println(xs[len(xs)-1:]) // * only the last one

	xi := []int{42, 43, 44, 45, 46, 47}
	fmt.Println("xi:", xi)

	// NOTE: delete from a slice
	xi = append(xi[:2], xi[3:]...) // * use ... to change slice to parameters of a function
	fmt.Println("xi:", xi)

	// NOTE: We can use make for slices, maps and channels
	// * creating slice using make
	xx := make([]int, 0, 10)
	// * 0  -> zero-value items count
	// * 10 -> default location space -> to prevent allocation overhead as we allocate whole new space when we append to a slice that it's capacity is fulled
	fmt.Println(xx)
	fmt.Println("len:", len(xx))
	fmt.Println("cap:", cap(xx)) // output: 10

	xx = append(xx, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("len:", len(xx))
	fmt.Println("cap:", cap(xx)) // output: 10

	xx = append(xx, 10, 11, 12, 13)
	fmt.Println("len:", len(xx))
	fmt.Println("cap:", cap(xx)) // output: 20
	// NOTE: Every time the slice reaches it's capacity the capacity will be increased by the first capacity

	// * multi-dimensional slice
	amir := []string{"Amir", "Zare", "27"}
	faeze := []string{"Faeze", "Arefi", "24"}
	fmt.Println(amir)
	fmt.Println(faeze)

	md := [][]string{amir, faeze}
	fmt.Println(md)
	fmt.Printf("%T \t %#v\n", md, md)
	fmt.Printf("%T \t %v\n", md, md)
}
