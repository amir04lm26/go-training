package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(15 * time.Second)
}

// * gen is a broken generator that will leak a goroutine
func gen() <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		// * continues infinitely
		for {
			ch <- n
			n++
		}
	}()

	return ch
}
