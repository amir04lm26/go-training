package main

import (
	"fmt"
	"time"
)

// NOTE: Timers and Tickers
/*
	In Go, you can create timers and tickers using the time package.
	Timers are used for one-time events, while tickers are for recurring events.
*/

// NOTE: A timer is a way to schedule something to happen after a specific duration.

func main() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer...")
	<-timer.C // NOTE: Wait for the timer to finish
	fmt.Println("Timer fired")
}
