package main

import (
	"fmt"
	"time"
)

// NOTE: Select Statement
/*
	The select statement allows a goroutine to wait on multiple communication operations.
	It’s similar to a switch statement but for channels.
*/
/*
	•	The select statement allows a goroutine to listen on multiple channels simultaneously.
	•	It will proceed with the first channel that is ready to communicate.
*/

func selectExample() {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channelOne <- "One"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "Two"
	}()

	for integer := 0; integer < 2; integer++ {
		select {
		case msg1 := <-channelOne:
			fmt.Println("msg1:", msg1)
		case msg2 := <-channelTwo:
			fmt.Println("msg2:", msg2)
		}
	}
}

// NOTE: default-case
/*
	The default case in a select is executed immediately if none of the channels are ready.
	This is helpful to avoid blocking when you want to continue with some other logic.
*/
func defaultSelectExample() {
	var resultChannel chan string = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		resultChannel <- "message"
	}()

	// NOTE: Explanation
	/*	•	Since there is no message ready on ch immediately, the default case is executed.
		Without the default case, the select would block until a value is received.
	*/
	select {
	case msg := <-resultChannel:
		fmt.Println("msg:", msg)
	default:
		fmt.Println("No message received, continue working...")
	}
}

func main() {
	selectExample()
	defaultSelectExample()
}
