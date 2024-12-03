package main

import (
	"fmt"
	"time"
)

// NOTE: select Statement with Multiple Channels
/*
	The select statement is used to wait on multiple channel operations.
	It’s like a switch statement but for channels, allowing a program to handle whichever
	channel becomes ready first.
*/

// NOTE: Example: Using select to Handle Multiple Channels

// NOTE: Explanation:
/*
	•	select: The select statement listens on both ch1 and ch2.
		It will print the message from whichever channel sends first.
	•	Timeout: The time.After channel introduces a timeout. If no message is received
		within 3 seconds, the select block will handle the timeout case.

	In this case, the message from ch2 arrives first because ch2 simulates a faster goroutine.
*/

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Simulate a slow goroutine
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	// Simulate a faster goroutine
	go func() {
		time.Sleep(time.Second)
		channel2 <- "Message from channel 2"
	}()

	select {
	case message := <-channel1:
		fmt.Println(message)
	case message := <-channel2:
		fmt.Println(message)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: no messages received")
	}
}
