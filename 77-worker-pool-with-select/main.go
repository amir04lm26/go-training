package main

import (
	"fmt"
	"time"
)

// NOTE: Worker Pool Pattern Revisited with Select
/*
	You can combine the worker pool pattern with the select statement to handle channel
	operations in a more sophisticated way.
	This allows workers to process jobs and handle timeouts or cancellation.
*/

// NOTE: Example: Worker Pool with Select

// NOTE: Explanation:
/*
	•	select with channels: Workers wait for jobs or timeout using the select statement.
	•	Timeout handling: The time.After case prevents workers from blocking indefinitely
		when no jobs are available.
*/

func worker(id int, jobs <-chan int, result chan<- int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return // * Channel is closed
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second)
			result <- job * 2
		case <-time.After(2 * time.Second):
			fmt.Printf("Worker %d timed out\n", id)
		}
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, jobs, results)
	}

	for iterator := 1; iterator <= 5; iterator++ {
		jobs <- iterator
	}
	close(jobs)

	for iterator := 1; iterator <= 5; iterator++ {
		fmt.Println("Result:", <-results)
	}
}
