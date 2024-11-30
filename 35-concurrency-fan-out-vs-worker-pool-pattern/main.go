package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: Fan-Out vs Worker Pool
/*
	In Go, both fan-out and worker pool patterns are used for concurrent processing,
	but they serve different purposes and are structured differently.
	Here’s a breakdown of the two patterns and their differences:
*/

// NOTE: Fan-out Pattern
/*
	•	Concept: Fan-out involves spawning multiple goroutines to handle tasks concurrently.
	•	Purpose: It allows the parallel execution of independent tasks.
	•	How it Works:
	•	A single goroutine (or the main goroutine) creates multiple goroutines.
	•	Each goroutine performs a task independently.
	•	You “fan-out” tasks to multiple goroutines.
*/

// NOTE: When to Use Fan-out
/*
	•	When you have multiple independent tasks to execute in parallel.
	•	When tasks don’t need to share data or communicate with each other.
*/

// NOTE: Key Characteristics
/*
	•	No central control over task distribution.
	•	Useful when you just want concurrent execution.
*/

func fanOutWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker %d finished\n", id)
}

func fanOutExample() {
	var wg sync.WaitGroup
	numWOrkers := 5

	for iterator := 1; iterator <= numWOrkers; iterator++ {
		wg.Add(1)
		go fanOutWorker(iterator, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}

// NOTE: Worker Pool Pattern
/*
	•	Concept: A worker pool consists of a fixed number of goroutines (workers) that repeatedly
		pull tasks from a shared queue or channel and process them.
	•	Purpose: It limits concurrency by controlling the number of active workers, which prevents
		resource exhaustion.
	•	How it Works:
	•	A manager distributes tasks to workers via a task channel.
	•	Each worker listens on the task channel and processes tasks one by one.
*/

// NOTE: When to Use Worker Pool
/*
	•	When you have a large number of tasks and need to control the concurrency to prevent resource exhaustion.
	•	When the number of workers (goroutines) is fixed but the number of tasks varies.
*/

// NOTE: Key Characteristics
/*
	•	Workers pull tasks from a shared channel.
	•	Concurrency is limited by the number of workers.
*/

func workerPool(id int, jobsChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobsChannel {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func workerPoolExample() {
	const numWorkers = 3
	const numJobs = 5

	jobsChannel := make(chan int, numJobs)
	var wg sync.WaitGroup

	for iterator := 1; iterator <= numWorkers; iterator++ {
		wg.Add(1)
		go workerPool(iterator, jobsChannel, &wg)
	}

	for iterator := 1; iterator <= numJobs; iterator++ {
		jobsChannel <- iterator
	}
	close(jobsChannel)

	wg.Wait()
	fmt.Println("All jobs processed")
}

// NOTE: Differences between Fan-out and Worker Pool
/*
	Aspect				Fan-out	Worker Pool
	Structure			Multiple goroutines created for each task.					Fixed number of goroutines processing tasks from a shared channel.
	Concurrency 		control	No control over the number of goroutines.			Limits concurrency by controlling the number of workers.
	Task distribution	Tasks are independent.										Tasks are queued and distributed to workers.
	When to use			When tasks are independent and need to run concurrently.	When you need to control concurrency and process many tasks efficiently.
	Examples			Web scraping multiple URLs simultaneously.					Handling a high volume of HTTP requests with limited resources.

	In summary, fan-out is ideal when you need to run independent tasks concurrently, whereas a worker pool
	is useful when you have a lot of tasks but need to limit the number of concurrent goroutines.
*/

func main() {
	fanOutExample()
	workerPoolExample()
}
