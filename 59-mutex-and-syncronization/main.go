package main

import (
	"fmt"
	"sync"
)

// NOTE: Mutex and Synchronization
/*
	When multiple goroutines need to access shared data, mutexes are used to prevent race conditions.
	A mutex ensures that only one goroutine can access the critical section (shared resource) at a time.
*/

// NOTE: Example: Using Mutex for Safe Access to Shared Data

// NOTE: Explanation:
/*
	•	Mutex: The sync.Mutex is used to lock the critical section, which in this case is the increment operation on the
		shared variable counter. Without a mutex, the goroutines would race to modify counter, leading to a race condition.
	•	Lock/Unlock: mu.Lock() locks the critical section, ensuring that only one goroutine can modify counter at a time.
		mu.Unlock() releases the lock, allowing other goroutines to access the critical section.

	Without the mutex, the counter value would be inconsistent due to race conditions, but with the mutex, the value is correctly synchronized.
*/

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
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
