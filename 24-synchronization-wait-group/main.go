package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: Synchronization
/*
	Concurrency in Go can be powerful, but when multiple Goroutines access shared resources,
	we need to synchronize them to avoid race conditions. Go provides a few built-in
	synchronization mechanisms, including WaitGroup and Mutex from the sync package.
*/

// NOTE: sync.WaitGroup
/*
	A WaitGroup allows you to wait for a collection of Goroutines to finish executing.
	You can add Goroutines to the wait group, and then wait for all of them to complete
	before proceeding with the rest of your code.
*/

// NOTE: Explanation
/*
	•	wg.Add(1): Adds one to the wait group’s count, which represents a Goroutine that
		needs to finish.
	•	wg.Done(): Decrements the counter, signaling that the Goroutine has completed.
	•	wg.Wait(): Blocks until the counter goes to zero, meaning all Goroutines are done.
*/

// NOTE: Defer
/*
	The defer keyword in Go is used to ensure that a function call is executed after the
	surrounding function completes. The deferred function will execute just before the
	surrounding function returns, no matter how the function exits (normal return, error, panic, etc.).

	deferred functions won't run when os.Exit runs
	all deferred functions will run after panic, and panicking propagates to the top of call-stack
	in reverse order
*/

func worker(id int, wg *sync.WaitGroup) {
	// NOTE WaitGroup.Done
	/*
		Done decrements the [WaitGroup] counter by one.

		In this Go snippet, defer is used to schedule the execution of the wg.Done()
		function to run at the end of the worker function, regardless of how the
		function exits.
	*/
	defer wg.Done() // NOTE: Signal that this goroutine is done
	fmt.Printf("Worker %d starting\r\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\r\n", id)
}

func main() {
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 3; iterator++ {
		// NOTE: WaitGroup.Add
		/*
			Add adds delta, which may be negative, to the [WaitGroup] counter.
			If the counter becomes zero, all goroutines blocked on [WaitGroup.Wait] are released.
			If the counter goes negative, Add panics.

			Note that calls with a positive delta that occur when the counter is zero must happen
			before a Wait. Calls with a negative delta, or calls with a positive delta that start
			when the counter is greater than zero, may happen at any time.

			Typically this means the calls to Add should execute before the statement creating
			the goroutine or other event to be waited for.

			If a WaitGroup is reused to wait for several independent sets
			of events, new Add calls must happen after all previous Wait calls have returned.
		*/

		wg.Add(1)
		go worker(iterator, &wg)
	}

	// NOTE: WaitGroup.Wait
	/*
		Wait blocks until the [WaitGroup] counter is zero.
	*/
	// NOTE: Wait for all Goroutines to finish
	wg.Wait()
	fmt.Println("All workers finished")
}
