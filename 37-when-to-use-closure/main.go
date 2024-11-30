package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: When to use Closure
/*
	The two code snippets you posted look very similar, but there is a key difference in how the goroutines capture
	the loop variable i. This subtle difference can lead to bugs if not handled properly. Let’s dive into how each one behaves.
*/

// NOTE: Snippet 1
/*
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results, errChan)
		}(i)
	}
*/

// NOTE: How it works
/*
	•	The loop variable i is passed as a parameter to the goroutine function.
	•	Each iteration of the loop creates a new goroutine with its own copy of i passed as id.
	•	This ensures that each goroutine receives the correct value of i when it starts running, even if the loop continues.

	This version is correct and safe because each goroutine gets a different id value (1, 2, 3).
*/

// NOTE: Snippet 2
/*
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i, jobs, results, errChan)
		}()
	}
*/

// NOTE: How it works
/*
	•	The loop variable i is used directly inside the anonymous function.
	•	However, all goroutines capture the same i variable from the outer scope. Since goroutines may not start executing
		immediately, they might all see the final value of i (which will be 4 after the loop finishes).

	This can result in unexpected behavior, where all the goroutines might use the same (final) value of i = 4.
*/

// NOTE: Why does this difference matter?
/*
	•	Goroutines may not execute immediately after being started. If you use the outer loop variable directly, by the
		time the goroutine runs, the loop might have completed, leaving i with its final value (4).
	•	The first snippet avoids this issue by passing i as a parameter, creating an independent copy for each goroutine.
*/

// NOTE: Which version should you use?
/*
	•	Use the first version where the loop variable is passed as a parameter to the goroutine:

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results, errChan)
		}(i)
	}

	This ensures that each goroutine receives the correct id value.
*/

// NOTE: Summary
/*
	•	Snippet 1 (passing i as a parameter) ensures correct behavior by avoiding variable capture issues.
	•	Snippet 2 can lead to race conditions or incorrect values if the loop completes before the goroutines run.
*/

func worker(id *int) {
	fmt.Printf("Worker %d processing a job\n", *id)
}

// ! I don't see any race condition
func main() {
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			worker(&iterator)
		}()
	}

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			worker(&id)
		}(iterator * 10)
	}

	wg.Wait()
}
