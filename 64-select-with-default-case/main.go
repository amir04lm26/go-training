package main

import "fmt"

// NOTE: select with Default Case
/*
	You can add a default case to a select statement to handle situations when none of
	the channels are ready.
	The default case is executed immediately if no other cases are ready.
*/

// NOTE: Example: select with Default Case

// NOTE: Explanation:
/*
	•	default: If no messages are available on the ch channel, the default case is executed
		immediately, allowing the program to move on without blocking.
*/

func main() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received, proceeding without waiting.")
	}
}
