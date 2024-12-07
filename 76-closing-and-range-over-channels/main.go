package main

import "fmt"

// NOTE: Closing Channels and Range Over Channels
/*
	Closing a channel signals that no more values will be sent on that channel.
	This is important for managing the lifecycle of channels, especially when dealing
	with goroutines.
*/

// NOTE: Example: Closing Channels and Using range

// NOTE: Explanation:
/*
	•	close(jobs): Closing the channel jobs prevents any more values from being sent.
	•	range: The range keyword is used to receive values from a channel until it is closed.
		When jobs is closed, the loop terminates.
	•	done: The done channel is used to signal when the worker has finished processing all jobs.
*/

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// * Send jobs to the jobs channel
	go func() {
		for iterator := 1; iterator <= 5; iterator++ {
			jobs <- iterator
			fmt.Println("Sent job", iterator)
		}
		close(jobs) // * Close the channel when done sending
	}()

	// * Worker to receive jobs
	go func() {
		for job := range jobs {
			fmt.Println("Received job", job)
		}
		done <- true
	}()

	<-done
}
