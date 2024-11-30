package main

import (
	"fmt"
	"time"
)

// NOTE: Worker Pool
/*
	A worker pool is a set of goroutines that wait for tasks, process them, and then repeat.
	This pattern is widely used for distributing tasks efficiently across multiple workers.
*/

// NOTE: Explanation
/*
	•	Job: Represents a unit of work (could be anything like a task, computation, etc.).
	•	Worker: Processes jobs and sends results back through the results channel.
	•	Worker Pool: Multiple workers are processing jobs concurrently.
*/

type Job struct {
	id    int
	value int
}

func worker(id int, jobs <-chan Job, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.id)
		time.Sleep(500 * time.Millisecond)
		result <- (job.value * 2)
	}
}

func main() {
	numWOrkers := 3
	jobsChannel := make(chan Job, 10)
	resultsChannel := make(chan int, 10)

	for iterator := 1; iterator < numWOrkers; iterator++ {
		go worker(iterator, jobsChannel, resultsChannel)
	}

	for iterator := 1; iterator <= 5; iterator++ {
		jobsChannel <- Job{id: iterator, value: iterator * 10}
	}
	close(jobsChannel)

	for iterator := 1; iterator <= 5; iterator++ {
		fmt.Println("Result:", <-resultsChannel)
	}
}
