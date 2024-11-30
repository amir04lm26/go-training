package main

import (
	"fmt"
	"time"
)

// NOTE: Select Statement
/*
	The select statement in Go is used to handle multiple channels at once.
	It blocks until one of its cases can proceed, making it useful for working with
	multiple channels, timeouts, or controlling concurrent operations.
*/

// NOTE: Example: Using select with multiple channels
/*
	In this example, we’ll have two channels, ch1 and ch2. We’ll use the select statement
	to read from whichever channel is ready first.
*/

// NOTE: Explanation:
/*
	•	select: The select statement waits on multiple channel operations.
		It chooses whichever channel is ready and executes the corresponding case.
		If multiple channels are ready, it chooses one randomly.
	•	Goroutines: Two goroutines simulate the sending of data to the channels after a delay.
		The select statement will handle the reception of messages from the channels.

		In this example, ch2 sends its message first because it sleeps for only 1 second,
		while ch1 sleeps for 2 seconds.
*/

func selectWithMultipleChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for iterator := 0; iterator < 2; iterator++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

// NOTE: Select with Default Case
/*
	You can also add a default case in the select block, which is executed if no channels are ready.
	This prevents select from blocking.
*/

//NOTE: Explanation:
/*
	•	default: The default case in the select statement allows the program to continue
		doing other work if no channel is ready to proceed.
	•	Looping: The for loop continuously checks the channel using the select statement,
		allowing the program to do other work while waiting for a message.

	In this example, the program keeps printing “No message received, doing other work…”
	until the goroutine sends a message after 2 seconds.
*/

func selectWithDefault() {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Message from goroutine"
	}()

	for {
		select {
		case message := <-channel:
			fmt.Println("Received:", message)
			return
		default:
			fmt.Println("No message received, doing other work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	selectWithMultipleChannels()
	selectWithDefault()
}
