package main

import (
	"fmt"
	"time"
)

// NOTE: Concurrency Patterns
/*
	Some common concurrency patterns in Go:

	•	Fan-out, fan-in: Multiple goroutines work concurrently on different parts of the job,
		and the results are collected in a single channel.
	•	Worker pools: A pool of fixed-size workers, each processing a job from a shared queue.
	•	Timeouts: Using select to implement timeouts for operations.
*/

// NOTE: Timeout Example
func timeoutExample() {
	processChannel := make(chan string)

	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}(processChannel)

	select {
	case result := <-processChannel:
		fmt.Println("result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func main() {
	timeoutExample()
}
