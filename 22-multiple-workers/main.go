package main

import (
	"fmt"
	"time"
)

// NOTE: Example: Aggregating Results from Multiple Workers
/*
	In Go, you can specify whether a function is allowed to send or receive values on a channel
	using directional channels.

	Let’s break this down:

	•	Bidirectional channel (chan string): The function can both send and receive data.
	•	Send-only channel (chan<- string): The function can only send data into the channel.
	•	Receive-only channel (<-chan string): The function can only receive data from the channel.

	When to Use Each:

	•	chan string (Bidirectional): Use this when you need the flexibility to both send and receive
		data in the same function. For example, if a function is both a producer and consumer of data.
	•	chan<- string (Send-Only): Use this when the function only needs to send data to the channel
		and should not be allowed to receive from it. This improves the clarity of your code, making
		it clear that this function is a producer of data.
	•	<-chan string (Read-Only): Use this when you want to make it clear that a function should
		only consume data from a channel (e.g., in worker or listener routines), and it should not
		send any data.
*/

// Worker function simulating work and sending result to a channel
func worker(id int, resultChannel chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println("After sleep:", id)
	resultChannel <- fmt.Sprintf("Worker %d finished", id)
}

func withTimeoutThatNotWorking() {
	fmt.Println("#withTimeoutThatNotWorking")
	resultChannel := make(chan string, 3)

	// Start 3 workers
	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, resultChannel)
	}

	// Receive results using select
	// NOTE: Is not the same as sequential select statement
	for iterator := 0; iterator < 3; iterator++ {
		select {
		case res := <-resultChannel:
			fmt.Println(res)
		// ! Why not the second response is also available?
		// * Reading from unBuffered channels blocks when the buffer is empty.
		// * Also, reading from unBuffered channels always block.
		// * So every select waits for it's response to receive and each
		// * select statement has its own timer which starts after the blocked time.
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: Worker took too long")
		}
	}
}

func withTimeoutWhichIsWorking() {
	fmt.Println("\r\n#withTimeoutWhichIsWorking")
	resultChannel := make(chan string, 3)

	for iterator := 1; iterator <= 3; iterator++ {
		go worker(iterator, resultChannel)
	}

	timeout := time.After(2 * time.Second)

	for iterator := 0; iterator < 3; iterator++ {
		select {
		case res := <-resultChannel:
			fmt.Println("res:", res)
		case <-timeout:
			fmt.Println("Timeout: Worker took too long")
		}
	}
}

func main() {
	withTimeoutThatNotWorking()
	withTimeoutWhichIsWorking()
}
