package main

import (
	"fmt"
	"time"
)

// NOTE: Concurrency Patterns in Go

/*
	Go has a rich set of concurrency tools, and we’ve already touched on goroutines and channels.
	Let’s look at more advanced concurrency patterns used in Go programs, including fan-in, fan-out,
	and worker pools.
*/

// NOTE: Fan-In pattern
/*
	The fan-in pattern merges multiple input channels into a single output channel. This pattern is
	useful when you want to gather results from multiple concurrent tasks into one stream of output.
*/

// NOTE: Explanation:
/*
	•	Producer: Simulates two goroutines (Producer 1 and Producer 2) sending data through channels.
	•	Fan-In: A function that merges multiple input channels (ch1, ch2) into a single output channel.
	•	Select Statement: The select block allows the program to receive from multiple channels concurrently.
*/

func producer(name string, channel chan<- string) {
	for iterator := 0; iterator < 3; iterator++ {
		channel <- fmt.Sprintf("%s: message %d", name, iterator)
		time.Sleep(500 * time.Millisecond)
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-input1:
				channel <- msg
			case msg := <-input2:
				channel <- msg
			}
		}
	}()

	return channel
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go producer("Producer 1", channel1)
	go producer("Producer 2", channel2)

	merged := fanIn(channel1, channel2)

	for iterator := 0; iterator < 6; iterator++ {
		fmt.Println(<-merged)
	}
}
