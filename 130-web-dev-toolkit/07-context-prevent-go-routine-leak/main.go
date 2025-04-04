package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	time.Sleep(15 * time.Second)
}

// * gen is a broken generator that will leak a goroutine
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return // * avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}
