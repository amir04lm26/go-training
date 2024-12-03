package main

import (
	"fmt"
	"time"
)

// NOTE: Rate Limiting with time.Tick
/*
	Rate limiting is used to control how frequently an operation is performed. In Go,
	the time.Tick function can create a ticker that generates events at a fixed interval.
*/

// NOTE: Example: Basic Rate Limiting

// NOTE: Explanation:
/*
	•	time.Tick: The time.Tick function returns a channel that sends the current time every
		500 milliseconds.
	•	Rate limiting: Each iteration of the loop is delayed until the ticker allows the next
		operation.
*/

func main() {
	limiter := time.Tick(500 * time.Millisecond)

	for iterator := 1; iterator <= 5; iterator++ {
		<-limiter
		fmt.Println("Request", iterator, "processed at", time.Now())
	}
}
