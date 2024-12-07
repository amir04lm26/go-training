package main

import (
	"context"
	"fmt"
	"time"
)

// NOTE: Contexts in Go
/*
	The context package in Go is designed to carry deadlines, cancellation signals, and other
	request-scoped values across API boundaries and goroutines. Contexts are useful for
	controlling the lifetime of processes, especially when multiple goroutines need to be aware
	of the same “cancellation” or “timeout” signals.
*/

// NOTE: Key functions in the context package:
/*
	•	context.Background(): Returns an empty context. It is typically used as the root of the
		context tree.
	•	context.WithCancel(): Returns a copy of the parent context with a new Done channel that
		is closed when the cancel function is called.
	•	context.WithTimeout(): Returns a copy of the parent context that is automatically
		canceled after a specified timeout.
	•	context.WithDeadline(): Similar to WithTimeout, but allows you to set a specific
		deadline (absolute time).
	•	context.WithValue(): Returns a copy of the parent context with an additional key-value pair.
		It’s useful for passing request-scoped data.
*/

// NOTE: Example: Simple context.WithCancel

// NOTE: Explanation:
/*
	•	context.WithCancel: The cancel function is called after 2 seconds, signaling the doWork
		function to stop by closing the Done channel.
	•	ctx.Done(): The Done channel is used to detect when the context has been canceled.
*/

func doWorkWithCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func contextWithCancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWorkWithCancel(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

// NOTE: Context with Timeout
/*
	context.WithTimeout is useful when you want to automatically cancel a context after
	a specific period of time.
*/

// NOTE: Example: context.WithTimeout

// NOTE: Explanation:
/*
	•	context.WithTimeout: The context will be canceled automatically after 1 second.
	•	ctx.Done(): The doWorkWithTimeout function listens for either the completion of the
		work or the cancellation signal.

	The work was canceled after 1 second because the context timed out, even though the task
	would have taken 2 seconds to complete.
*/

func doWorkWithTimeout(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work canceled:", ctx.Err())
	}
}

func contextWithTimeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWorkWithTimeout(ctx)

	time.Sleep(2 * time.Second)
}

// NOTE: Context with Value
/*
	context.WithValue allows you to associate key-value pairs with a context, which is
	useful for passing metadata like authentication tokens or request IDs across API
	boundaries.
*/

// NOTE: Example: context.WithValue

// NOTE: Explanation:
/*
	•	context.WithValue: Adds a key-value pair (userID: 42) to the context.
	•	ctx.Value(): Retrieves the value associated with the key "userID" in the
		processRequest function.
*/

// * Define a custom type for the context key
type contextKey string

const userIdKey contextKey = "userID"

func processRequest(ctx context.Context) {
	if userID, ok := ctx.Value(userIdKey).(int); ok {
		fmt.Println("Processed request for user:", userID)
	} else {
		fmt.Println("User ID not found")
	}
}

func contextWithValueExample() {
	ctx := context.WithValue(context.Background(), userIdKey, 42)
	processRequest(ctx)
}

func main() {
	contextWithCancelExample()
	contextWithTimeoutExample()
	contextWithValueExample()
}
