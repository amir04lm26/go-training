package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// NOTE: Atomic Operations
/*
	In the context of avoiding race conditions in concurrent Go programs, using atomic operations
	from the sync/atomic package is a good solution for managing shared resources between goroutines
	without the need for explicit locking (like sync.Mutex). Atomic operations are efficient and
	ensure memory consistency across multiple goroutines when updating shared variables.

	This is a practical and efficient way to handle concurrent updates to shared state in Go when
	the operations on the shared state are simple (e.g., incrementing, setting a flag).
*/

// NOTE: Explanation of key components
/*
1.	sync/atomic.AddInt64:
	•	This is an atomic operation that increments the shared variable counter by 1 in a thread-safe
		manner.
	•	It ensures that even though multiple goroutines access and modify counter concurrently, they
		do so without causing race conditions.
2.	int64 Counter:
	•	The counter variable is an int64, which is compatible with the atomic package’s functions
		like AddInt64.
	•	Passing a pointer to counter (&counter) allows the atomic.AddInt64 function to modify the
		original variable.
3.	Goroutines:
	•	Each worker goroutine increments the counter 1000 times.
	•	Using atomic operations ensures that these increments are safe, and no updates to counter are
		lost.
4.	sync.WaitGroup:
	•	This is used to wait for all the worker goroutines to finish before printing the final value
		of the counter.
	•	wg.Add(1) increments the wait group counter, and wg.Done() decrements it, signaling when a
		goroutine is done.
*/

// NOTE: Benefits of Atomic Operations
/*
	•	No Locking Overhead: Unlike sync.Mutex, atomic operations don’t require locking, so they’re
		generally faster for simple cases like incrementing counters.
	•	Race Condition Prevention: Atomic operations guarantee that increments or decrements are
		performed atomically, preventing race conditions that can occur when multiple goroutines
		try to update a shared variable simultaneously.
*/

func worker(id int, counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for iterator := 0; iterator < 1000; iterator++ {
		atomic.AddInt64(counter, 1) // NOTE: Atomically increment the counter
		runtime.Gosched()           // yields the processor
		// fmt.Println("counter:", atomic.LoadInt64(counter))
	}

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var counter int64
	var wg sync.WaitGroup

	workersCount := 5

	for iterator := 0; iterator < workersCount; iterator++ {
		wg.Add(1)
		go worker(iterator+1, &counter, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
