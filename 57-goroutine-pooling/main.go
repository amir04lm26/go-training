package main

import (
	"fmt"
	"time"
)

// NOTE: Goroutine Pooling
/*
	Goroutine pooling is a technique used to control the number of goroutines executing concurrently,
	which helps in managing resources and preventing excessive resource usage.
*/

// NOTE: Example: Worker Pool
/*
	In this example, we’ll implement a simple worker pool where we limit the number of goroutines that process tasks from a job queue.
*/

// NOTE: Explanation:
/*

	•	Worker pool: We start a fixed number of workers (3 in this example), each running as a goroutine.
		These workers read from the jobs channel and process the tasks concurrently.
	•	Jobs queue: The jobs are sent to the jobs channel, and each worker picks up jobs from this queue.
	•	Results: After processing the jobs, the workers send the results to the results channel,
		which the main function collects and prints.

	In this example, three workers process five jobs concurrently. When a worker finishes a job,
	it picks up the next one from the queue.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep((time.Second))
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	numWOrkers := 3
	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for iterator := 1; iterator <= numWOrkers; iterator++ {
		go worker(iterator, jobs, results)
	}

	for iterator := 1; iterator <= numJobs; iterator++ {
		jobs <- iterator
	}
	close(jobs)

	for iterator := 1; iterator <= numJobs; iterator++ {
		fmt.Println("Results", <-results)
	}

	// !NOTE: deadlock! because the channel is not closed
	// for result := range results {
	// 	fmt.Println("Result:", result)
	// }
}
