package main

import (
	"fmt"
	"time"
)

// NOTE: Buffered Channels and Rate Limiting
/*
	You can use buffered channels to allow a burst of events while still limiting the overall rate.
*/

// NOTE: Example: Bursty Rate Limiting

// NOTE: Explanation:
/*
	•	burstyLimiter: This buffered channel allows a burst of up to 3 requests at once.
		The limiter goroutine refills the channel every 500 milliseconds, limiting the rate after
		the initial burst.

		The first three requests are processed immediately due to the burst limit, and the subsequent
		requests are processed at the rate of one every 500 milliseconds.
*/

func main() {
	limiter := time.Tick(500 * time.Millisecond)
	burstyLimiter := make(chan time.Time, 3)

	// * Fill up the burstyLimiter channel, allowing a burst of 3 events
	for iterator := 0; iterator < 3; iterator++ {
		burstyLimiter <- time.Now()
	}

	// * Start a goroutine to add a new value to burstyLimiter every 500ms
	go func() {
		for timer := range limiter {
			burstyLimiter <- timer
		}
	}()

	for iterator := 1; iterator <= 5; iterator++ {
		<-burstyLimiter
		fmt.Println("Request", iterator, "processed at", time.Now())
	}
}
