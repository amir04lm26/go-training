package main

import (
	"fmt"
	"sync"
)

// NOTE: Race Condition
/*
	A race condition occurs when two or more Goroutines or threads access shared data concurrently,
	and the final outcome depends on the order in which these accesses occur.
	This leads to unpredictable behavior, as the results can vary depending on how the Go runtime
	schedules the Goroutines. Race conditions are a common problem in concurrent programming and
	can cause bugs that are difficult to reproduce and diagnose.
*/

// NOTE: Example explanation
/*
	Imagine two Goroutines trying to update the same variable at the same time.
	Depending on which Goroutine performs its action first, the variable might
	end up with different values.
	This non-deterministic behavior is what makes race conditions tricky.
*/

// NOTE: What's happening in the example:
/*
	•	The program starts 3 Goroutines, each incrementing the global counter variable 1000 times.
	•	Because these Goroutines are running concurrently, they may access the shared counter
		variable at the same time.
	•	Expected Output: You might expect the final value of counter to be 3000 (3 Goroutines × 1000
		increments each).
	•	Actual Output: The final value of counter will often be less than 3000 due to a race condition.
*/

// NOTE: How to prevent race conditions
/*
	To avoid race conditions, you need to control access to shared data. The most common solution
	is to use synchronization mechanisms, such as mutexes or atomic operations.
*/

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for iterator := 0; iterator < 1000; iterator++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

	for iterator := 0; iterator < 3; iterator++ {
		wg.Add(1)
		go increment(&wg) // NOTE: Run multiple Goroutines that increment the same variable
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
