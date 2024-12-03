package main

import (
	"fmt"
	"sync"
)

// NOTE: Mutexes in Go
/*
	A mutex (mutual exclusion) is a synchronization primitive used to ensure that only one goroutine
	can access a critical section of code at a time, preventing race conditions.

	In Go, the sync.Mutex type provides two methods:
		•	Lock(): Blocks if the mutex is already locked by another goroutine. Once unlocked,
			it locks the mutex for the calling goroutine.
		•	Unlock(): Unlocks the mutex so another goroutine can acquire the lock.
*/

// NOTE: Example: Using Mutex to Avoid Race Conditions

// NOTE: Explanation:
/*
	•	sync.Mutex: The mu field is a mutex that protects access to the value field.
	•	Increment method: This method locks the mutex before modifying the value and unlocks it
		after the operation is done.
	•	main function: The program starts 1000 goroutines, each incrementing the counter,
		and the WaitGroup ensures that the program waits for all goroutines to finish.

	Without the mutex, race conditions would cause the final counter value to be inconsistent.
*/

type Counter struct {
	mu    sync.Mutex
	value int
}

func (counter *Counter) Increment() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value++
}

func (counter *Counter) Value() int {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for iterator := 0; iterator < 1000; iterator++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
