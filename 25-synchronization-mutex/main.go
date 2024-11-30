package main

import (
	"fmt"
	"sync"
)

// NOTE: Sync.Mutex
/*
	A Mutex is used to lock and unlock critical sections of code to prevent
	race conditions when multiple Goroutines are accessing shared variables.
*/

// NOTE: Explanation
/*
	•	mu.Lock(): Locks the critical section of code, ensuring no other Goroutine
		can access the shared resource (counter) until the lock is released.
	•	mu.Unlock(): Unlocks the critical section, allowing other Goroutines to proceed.
	•	This prevents race conditions, where multiple Goroutines try to modify the same
		variable concurrently, leading to incorrect results.
*/

var counter int
var mu sync.Mutex

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()   // Lock the critical section
	counter++   // Access the shared variable
	mu.Unlock() // Unlock the critical section
}

func main() {
	fmt.Println("Counter:", counter) // NOTE: The default value of an int type is zero

	var wg sync.WaitGroup

	for iterator := 0; iterator < 1000; iterator++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Print("Final counter value:", counter, "\n") // NOTE: Print doesn't add '\n' at the end of the printing value
}
