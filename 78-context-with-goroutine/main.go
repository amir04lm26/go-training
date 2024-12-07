package main

import (
	"context"
	"fmt"
	"time"
)

// NOTE: Context with Goroutines
/*
	The context package is often used to cancel goroutines that are no longer needed.
*/

// NOTE: Example: Canceling Goroutines with Context

func doWork(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d canceled\n", id)
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for iterator := 1; iterator <= 3; iterator++ {
		go doWork(ctx, iterator)
	}

	time.Sleep(2 * time.Second)
	cancel() // * Cancel all workers

	time.Sleep(time.Second) // * Give goroutines time to finish
}
