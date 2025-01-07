package main

import (
	"fmt"
	"runtime"
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
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	go say("Hello")
	say("World")
}
