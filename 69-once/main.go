package main

import (
	"fmt"
	"sync"
)

// NOTE: Once: Ensuring One-Time Initialization
/*
	The sync.Once type ensures that a piece of code is executed only once, no matter how many times
	it’s called. This is often used for one-time initialization of resources.
*/

// NOTE: Example: Using sync.Once for Initialization

// NOTE: Explanation:
/*
	•	sync.Once: The once.Do method ensures that the initialize function is called only once,
		regardless of how many goroutines invoke it.

	The initialization code is run only once, even though multiple workers attempt to invoke it.
*/

var once sync.Once

func initialize() {
	fmt.Println("Initialization preformed")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize)
	fmt.Println("Worker running")
}

func main() {
	var wg sync.WaitGroup

	for iterator := 0; iterator < 3; iterator++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
}
