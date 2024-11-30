package main

import (
	"context"
	"fmt"
	"time"
)

// NOTE: The Context Package
/*
	The context package provides a way to pass deadlines, cancellations, and request-scoped
	values across Goroutines. This is crucial when building concurrent programs that might
	need to be canceled or timed out.

	The context package is essential for managing long-running Goroutines and processes,
	especially when you need to cancel them based on conditions like timeouts or user input.
*/

// NOTE: The Purpose of context.WithTimeout
/*
	•	context.WithTimeout creates a context that automatically gets canceled after a
		specified duration.
	•	You typically use this when you want to limit the amount of time a function
		(like a database query, HTTP request, or other I/O-bound operations) can take to
		prevent hanging or over-consuming resources.
*/

// NOTE: Explanation
/*
	•	context.Background(): This returns an empty context, which is the root context
		typically used at the top of your call tree.
	•	context.WithTimeout(): This function takes a parent context (context.Background())
		and a timeout duration (in this case, 2 seconds).
	•	ctx: The new context returned, which will be automatically canceled after 2 seconds.
	•	cancel: A function that can be called manually to cancel the context earlier. It’s
		always a good practice to defer the cancel() call to ensure proper cleanup (even if
		timeout occurs).
	•	defer cancel(): This ensures that the resources associated with the context (like timers)
		are freed up when the function returns, preventing memory leaks.
	•	time.After(1 * time.Second): This simulates a task that takes 1 second to complete.
		The select statement waits for this case to finish.
	•	ctx.Done(): This channel is closed when the context is canceled (either after 2 seconds
		or if cancel() is manually called).
	•	If the task finishes within the timeout (in this case, after 1 second), you’ll see
		"Task completed".
	•	If the task takes longer than the timeout (which it doesn’t here), the case with ctx.Done()
		will trigger, printing "Task canceled:" along with the error (ctx.Err()), which will be
		"context deadline exceeded".
*/

// NOTE: Why context.Context Is Passed as the First Argument
/*
	In Go, it’s a convention to pass context.Context as the first argument to functions that use it. This is because:

	1.	Consistency: By always placing the context as the first argument, it’s easy to understand at a glance which functions support cancellation and timeouts.
	2.	Readability: It makes the function signatures uniform and easy to spot when looking at function definitions or calling functions.
*/

// NOTE: context Package Key Functions
/*
	•	context.Background(): A base context that doesn’t get canceled or timeout.
	•	context.WithCancel(): Creates a context that can be canceled manually.
	•	context.WithTimeout(): Creates a context that is automatically canceled after a certain time.
	•	context.WithDeadline(): Creates a context that is canceled at a specific time (absolute deadline).
	•	ctx.Done(): A channel that gets closed when the context is canceled.
	•	ctx.Err(): Returns an error indicating why the context was canceled (context.Canceled or context.DeadlineExceeded).
*/

// NOTE: context.WithTimeout
func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

// NOTE: context.WithCancel - longRunningTask
func longRunningTask(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d cancelled\n", id)
		default:
			fmt.Printf("Task %d is running...\n", id)
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}
}

// NOTE: context.WithCancel
func withCancel() {
	// NOTE: Long Running Task
	longCtx, longCancel := context.WithCancel(context.Background())

	go longRunningTask(longCtx, 1)
	go longRunningTask(longCtx, 2)

	time.Sleep(2 * time.Second)

	longCancel()

	time.Sleep(1 * time.Second)
}

// NOTE: context.WithDeadline - taskWithDeadline
func taskWithDeadline(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d is canceled duo to deadline\n", id)
			return
		default:
			fmt.Printf("Task %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// NOTE: context.WithDeadline
func withDeadline() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go taskWithDeadline(ctx, 1)

	time.Sleep(3 * time.Second)
}

// NOTE: Passing Context
func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled!")
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func withPassingContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(4 * time.Second)
}

func main() {
	withTimeout()
	withCancel()
	withDeadline()
	withPassingContext()
}
