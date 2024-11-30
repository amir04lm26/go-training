package main

import (
	"fmt"
	"time"
)

// NOTE: A ticker is used when you want an action to be performed at regular intervals.

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for tick := range ticker.C {
			fmt.Println("Tick at", tick)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
