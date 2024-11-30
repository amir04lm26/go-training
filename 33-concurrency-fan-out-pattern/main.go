package main

import (
	"fmt"
	"time"
)

// NOTE: Explanation:
/*
	•	Worker: The worker function simulates a worker that reads from the jobs channel.
	•	Fan-Out: The fan-out pattern is implemented by launching multiple workers, each reading from the same jobs channel.
*/

func worker(id int, jobsChannel <-chan int) {
	for job := range jobsChannel {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	jobsChannel := make(chan int, 10)

	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, jobsChannel)
	}

	// for job := 0; job <= 9; job++ {
	for job := 0; job <= 18; job++ {
		jobsChannel <- job
	}
	close(jobsChannel)

	// NOTE: Wait for the jobs to be done (if job <= 18)
	// time.Sleep(2 * time.Second)
}
