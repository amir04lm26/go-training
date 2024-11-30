package main

import (
	"context"
	"fmt"
	"time"
)

// NOTE: Handling Context in Pipelines
/*
	When building long-running pipelines, it’s important to allow for cancellation and timeouts.
	This is where the context package comes in handy, especially for controlling goroutines and
	ensuring resources are cleaned up properly.
*/

// NOTE: Explanation
/*
	•	context.Context: Used to control the lifecycle of goroutines in the pipeline.
		If the context is canceled (in this case, due to a timeout), the pipeline is shut down.
	•	ctx.Done(): A channel that signals when the context is canceled, allowing each stage to
		stop processing.
*/

func generateWithContext(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			// time.Sleep(750 * time.Millisecond) // NOTE: To simulate the cancellation
			select {
			case <-ctx.Done():
				return
			case out <- num:
			}
		}
	}()
	return out
}

func squareWithContext(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-ctx.Done():
				return
			case out <- num * num:
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	numbers := generateWithContext(ctx, 2, 3, 4, 5)
	squared := squareWithContext(ctx, numbers)

	for res := range squared {
		fmt.Println(res)
	}
}
