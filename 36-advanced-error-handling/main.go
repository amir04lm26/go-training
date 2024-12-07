package main

import (
	"errors"
	"fmt"
	"sync"
)

// NOTE: Advanced Error Handling
/*
	Error handling in Go is simple and explicit, but there are several patterns that improve the readability and robustness of code.
*/

// NOTE: Returning Errors
/*
	Returning errors is the most common form of error handling in Go.
*/

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

// NOTE: Explanation
/*
	•	Errors in Go: Go uses the error type to represent errors. You can return both a result
		and an error, and handle the error appropriately in the calling code.
	•	errors.New: Creates a new error when needed, for example, when dividing by zero.
*/

func returningErrorsExample() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// NOTE: Custom Error Types
/*
	You can define custom error types to provide more information or structure for your errors.
	You can define custom error types by implementing the Error() method.
*/

type DivideError struct {
	Divisor  int
	Dividend int
	Message  string
}

// NOTE: Explanation
/*
	The Error() method returns a string representation of the error. Any type that implements the
	Error() string method automatically satisfies the error interface in Go.

	type error interface {
		Error() string
	}
*/
func (err *DivideError) Error() string {
	return fmt.Sprintf("Cannot divide %d by %d: %s", err.Dividend, err.Divisor, err.Message)
}

func divideCustomError(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, &DivideError{Divisor: num2, Dividend: num1, Message: "division by zero"}
	}
	return num1 / num2, nil
}

// NOTE: Explanation
/*
	•	Custom Error Type: DivideError is a custom error type that implements the error interface by defining an Error() method.
	•	Rich Error Information: You can store additional information (like the divisor and dividend) inside the error, making debugging easier.
*/

func customErrorTypeExample() {
	_, err := divideCustomError(4, 0)
	if err != nil {
		fmt.Println(err)
	}
}

// NOTE: Concurrency and Error Handling
/*
	When working with concurrency, error handling becomes more complicated. You need to ensure that
	errors from multiple goroutines are handled properly.
*/

func worker(id int, jobsChan <-chan int, resultsChan chan<- int, errChan chan<- error) {
	for job := range jobsChan {
		if job == 3 {
			errChan <- fmt.Errorf("job %d caused an error", job)
			return
		}
		fmt.Printf("Worker %d processing job %d\n", id, job)
		resultsChan <- job * 2
	}
}

func concurrentErrorHandlingExample() {
	jobsChan := make(chan int, 5)
	resultChan := make(chan int, 5)
	errChan := make(chan error, 1) // ! More than one error will cause a deadlock!
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobsChan, resultChan, errChan)
		}(iterator)
	}

	for iterator := 1; iterator <= 5; iterator++ {
		jobsChan <- iterator
	}
	close(jobsChan)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	for result := range resultChan {
		fmt.Println("Result:", result)
	}

	if err := <-errChan; err != nil {
		fmt.Println("Error:", err)
	}

	if err := <-errChan; err != nil {
		fmt.Println("Error:", err)
	}
}

/// NOTE: Explanation
/*
	•	Error Channel: errChan is used to capture errors from goroutines. It’s important to ensure that errors from
		multiple concurrent tasks are handled properly.
	•	Sync and Channels: The sync.WaitGroup ensures that all goroutines complete, while channels ensure error handling.
*/

// NOTE: Wrapping and Unwrapping Errors
/*
	Go 1.13 introduced error wrapping to add context to errors.
*/

// NOTE: Explanation
/*
	•	Wrapping Error: The fmt.Errorf function wraps the original error (division error) with
		additional context (cannot divide by zero).
	•	Unwrapping Errors: You can later inspect or “unwrap” these errors to check the original cause.
*/

func divideWithErrorWrapper(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero: %w", errors.New("division error"))
	}
	return num1 / num2, nil
}

func wrappingAndUnwrappingErrors() {
	_, err := divideWithErrorWrapper(10, 0)
	if err != nil {
		fmt.Println("wrapped error:", err)
	}
}

// NOTE: Error Handling Best Practices
/*
	•	Check errors immediately: Always check the error returned from functions and handle them
		as close as possible to where they occur.
	•	Wrap errors for context: Use fmt.Errorf to wrap errors with more context, especially if
		you are passing the error up the call stack.
	•	Return early: If an error occurs, return immediately rather than nesting logic.
*/

func main() {
	returningErrorsExample()
	customErrorTypeExample()
	concurrentErrorHandlingExample()
	wrappingAndUnwrappingErrors()
}
