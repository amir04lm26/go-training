package main

import (
	"fmt"
	"time"
)

// NOTE: Rate Limiting with Goroutines
/*
	Rate limiting controls how frequently actions are allowed to happen.
	Go provides built-in mechanisms for rate limiting using channels and time.Ticker.
*/

// NOTE: Explanation:
/*

	•	Ticker: The time.Ticker allows you to send events at regular intervals (500ms in this case).
		Each request is processed only after receiving from the ticker.
	•	Rate Limiting: The requests are processed at a controlled rate, with a 500ms delay between each.

	In this case, each request is processed 500ms apart, demonstrating basic rate limiting using time.Ticker.
*/

// NOTE: Example: Rate Limiting with time.Ticker

func main() {
	requestsChan := make(chan int, 5)
	for iterator := 1; iterator <= 5; iterator++ {
		requestsChan <- iterator
	}
	close(requestsChan)

	limiter := time.NewTicker(500 * time.Millisecond)

	for request := range requestsChan {
		<-limiter.C
		fmt.Println("Processed request", request, time.Now())
	}
}
