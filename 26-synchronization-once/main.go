package main

import (
	"fmt"
	"sync"
)

// NOTE: Sync.Once
/*
	The sync.Once construct allows a piece of code to be executed only once,
	even if it is called from multiple Goroutines. This is useful for things
	like initializing a resource or loading configuration.
*/

// NOTE: Explanation
/*
	once.Do(initialize): Ensures that initialize() is only called once,
	regardless of how many Goroutines attempt to call it. This is useful
	for ensuring one-time operations, like setting up a database connection.
*/

var once sync.Once

func initialize() {
	fmt.Println("Initialization preformed")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize) // NOTE: Ensure initialization happens only once
	fmt.Println("Worker executing")
}

func main() {
	var wg sync.WaitGroup

	for iterator := 0; iterator < 3; iterator++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
