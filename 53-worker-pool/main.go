package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// NOTE: What is a Worker Pool?
/*
	A worker pool is a pattern where a fixed number of goroutines (workers) process tasks concurrently.
	Each worker receives tasks from a job queue (channel), processes them, and reports the results,
	often back to another channel for collecting results.

	This pattern is useful when you need to control the number of goroutines concurrently executing a task,
	especially in scenarios where creating too many goroutines might overwhelm system resources.
*/

// NOTE: Example: Worker Pool
/*
	Let’s create an example where workers process tasks concurrently. Here, we simulate a task as just
	sleeping for a random amount of time, and each worker processes a task from the job queue.
*/

type Job struct {
	id    int
	delay time.Duration
}

type Result struct {
	jobId    int
	duration time.Duration
}

func worker(id int, jobs <-chan Job, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.id)
		time.Sleep(job.delay)
		fmt.Printf("Worker %d finished job%d\n", id, job.id)
		result <- Result{jobId: job.id, duration: job.delay}
	}
}

// NOTE: Explanation of the Example
/*
	•	Job struct: Represents a task that has an ID and a delay. The delay is used to simulate a time-consuming task.
	•	Result struct: Holds the result of a processed job, including the job ID and how long it took to complete.
	•	worker function: Each worker (goroutine) listens for jobs on the jobs channel. When a job is received, the worker
		processes the job (in this case, it sleeps for a given duration), then sends the result to the results channel.
	•	Channels:
	•	jobs: A channel that holds the tasks (jobs) that need to be processed.
	•	results: A channel that holds the results of the processed tasks.
	•	WaitGroup: We use a sync.WaitGroup to ensure that all workers complete before the program exits.
	•	Main logic:
	•	We first create the worker pool with numWorkers workers.
	•	Then, we send numJobs jobs to the jobs channel.
	•	After sending all jobs, we close the jobs channel to signal that no more jobs are coming.
	•	The program waits for all workers to finish processing (wg.Wait()), then closes the results channel.
	•	Finally, we read and print the results from the results channel.
*/

// NOTE: Key Concepts:
/*
	•	Buffered Channels: Both the jobs and results channels are buffered to ensure that we don’t block when sending jobs or results.
		Without buffering, a worker might block while waiting for space in the channel to send its result.
	•	Concurrency Control: We control the number of concurrent workers by limiting the number of goroutines (workers) processing jobs.
		This prevents creating too many goroutines, which can be inefficient.
*/

func main() {
	// NOTE: rand.Seed:
	/*
		rand.Seed is a function from Go's math/rand package.
		It sets the seed for the random number generator.
		The seed determines the sequence of random numbers produced by subsequent calls to rand functions.
		Without setting a seed, Go's math/rand generates the same sequence every time the program runs,
		which is useful for debugging but not for applications requiring true randomness.
	*/

	// NOTE: Why Use It?
	/*
		To make the output of random number generation less predictable by initializing the generator with a dynamic, high-precision value.
		This is critical in scenarios like:

		Games requiring unpredictable behavior.
		Simulations that need varied outcomes.
		Any application where true randomness is desirable.
	*/
	// rand.Seed(time.Now().UnixNano()) // ! deprecated

	// NOTE: Why Is rand.Seed Deprecated?
	/*
		Global State Issues: Using rand.Seed affects the global random number generator, which can lead to unintended
		interference when different parts of a program depend on randomness.
		Local Generators: The new approach promotes the use of locally scoped random number generators with rand.New,
		which are safer and more predictable for specific use cases.
	*/

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	const numWorkers = 3
	const numJobs = 5

	jobsChan := make(chan Job, numJobs)
	resultsChan := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for workerIterator := 1; workerIterator <= numWorkers; workerIterator++ {
		wg.Add(1)
		go worker(workerIterator, jobsChan, resultsChan, &wg)
	}

	for jobIterator := 1; jobIterator <= numJobs; jobIterator++ {
		// NOTE: rand.Intn(n):
		/*
			This generates a random integer between 0 and 2 (inclusive), because rand.Intn(n) generates a number in the range [0, n-1].
		*/
		job := Job{id: jobIterator, delay: time.Duration(random.Intn(3)+1) * time.Second}
		jobsChan <- job
	}
	close(jobsChan)

	wg.Wait()
	close(resultsChan)

	for result := range resultsChan {
		fmt.Printf("Job %d took %v to complete\n", result.jobId, result.duration)
	}
}
