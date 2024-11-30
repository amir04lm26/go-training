package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for iterator := 0; iterator < 10; iterator++ {
		fmt.Println(
			pos(iterator),
			neg(-2*iterator),
		)
	}
}
