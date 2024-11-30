package main

import (
	"fmt"
	"time"
)

// NOTE: Fan-In and Fan-Out Pattern
/*
	Fan-in and Fan-out are common concurrency patterns used in Go to manage work distribution.
		•	Fan-Out: Distribute work across multiple goroutines.
		•	Fan-In: Aggregate the results from multiple goroutines.
*/

// NOTE: Fan-Out Example: Distribute Work to Multiple Workers
/*
	Let’s say we have multiple workers that process jobs concurrently, and each worker runs
	as a goroutine.
*/

// NOTE: Explanation:
/*
	•	Fan-Out: We have 3 workers processing jobs concurrently. Each worker is a goroutine that processes jobs from the jobs channel.
	•	Fan-In: The results from the workers are collected into the results channel and printed in the main function.
*/

// NOTE: Key Points:
/*
	•	Fan-Out: We start multiple workers to process jobs concurrently.
	•	Fan-In: After the workers complete their tasks, we collect and process the results.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
		fmt.Printf("worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobsChan := make(chan int, 5)
	resultsChan := make(chan int, 5)

	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, jobsChan, resultsChan)
	}

	for iterator := 1; iterator <= 5; iterator++ {
		jobsChan <- iterator
	}
	close(jobsChan)

	for iterator := 1; iterator <= 5; iterator++ {
		result := <-resultsChan
		fmt.Println("Results:", result)
	}
}
