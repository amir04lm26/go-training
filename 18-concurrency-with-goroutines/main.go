package main

import (
	"fmt"
	"time"
)

func say(message string) {
	for integer := 0; integer < 3; integer++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

// NOTE: Use go before a function call to run it as a goroutine.
func main() {
	go say("Hello")
	say("World")
}
