package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// NOTE: Atomic Operations
/*
	In Go, atomic operations are used for low-level synchronization when multiple goroutines share a variable.
	The sync/atomic package provides functions to safely read, write, and modify integers and pointers without using mutexes.

	Atomic operations are faster than mutexes but are limited to basic operations like incrementing or comparing values.
	They are useful when you need very lightweight synchronization in performance-critical scenarios.
*/

// NOTE: Example: Atomic Counter

// NOTE: Explanation:
/*
	•	atomic.AddInt64: This function atomically adds 1 to the counter variable.
		It ensures that no two goroutines update the variable at the same time, avoiding race conditions.
	•	counter: counter is a global variable shared among multiple goroutines.
		Using atomic operations ensures that the increment operation is safe.

	Without atomic operations, the counter value could be inconsistent due to race conditions.
*/

var counter int64

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1) // * Atomically increment the counter
}

func main() {
	var wg sync.WaitGroup

	for iterator := 0; iterator < 1000; iterator++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
