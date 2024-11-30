package main

import (
	"fmt"
	"time"
)

// NOTE: Worker Pool
/*
	A worker pool is a concurrency pattern where multiple Goroutines (workers) process tasks
	from a queue (channel). This is useful when you have a lot of work to do and want to
	distribute it across multiple workers.
*/

// NOTE: Explanation
/*
	•	Workers: Goroutines that wait for jobs to be sent on the jobs channel, process them,
		and send the result back on the results channel.
	•	Job Queue: The jobs channel where tasks are sent to be processed by the workers.
	•	Results: Once the work is completed, the result is sent on the results channel.

	* 	This pattern is highly useful for parallelizing workloads efficiently, allowing you to
	* 	handle large workloads without overloading your system.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d.\r\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d.\r\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobs, results := make(chan int, 100), make(chan int, 100)

	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, jobs, results)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	for iterator := 1; iterator <= 5; iterator++ {
		fmt.Println("Result:", <-results)
	}
}
