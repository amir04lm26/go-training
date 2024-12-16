package main

import (
	"fmt"
	"math/rand/v2"
)

// NOTE: Control Flow
/*
	sequential
	conditional
	loop
*/

func main() {
	const base = 12

	if randNum := 2 * rand.Int64N(base); randNum >= base {
		fmt.Println("The condition is met so we will proceed")
		return
	} else {
		fmt.Println("We also access randNum here:", randNum)
	}

	// randNum // * We don't have a access to it here
	fmt.Println("The condition is not met")
}
