package main

import (
	"errors"
	"fmt"
)

// NOTE: Error Handling in Pipelines
/*
	One challenge with pipelines is handling errors. If an error occurs in one stage, you may
	want to signal the error to downstream stages. One approach is to pass both the result
	and the error through a channel
*/

// NOTE: Explanation
/*
	•	Error Struct: Instead of passing just the value through the pipeline, a struct (result)
		is used that holds both the value and any error.
	•	Error Propagation: If an error occurs, it’s propagated through the pipeline, and the
		pipeline is terminated early.
*/

type result struct {
	value int
	err   error
}

func generateWithError(nums ...int) <-chan result {
	out := make(chan result)
	go func() {
		for _, num := range nums {
			if num == 0 {
				out <- result{0, errors.New("cannot process zero")}
				close(out)
				return
			}
			out <- result{num, nil}
		}
		close(out)
	}()
	return out
}

func squareWithError(in <-chan result) <-chan result {
	out := make(chan result)
	go func() {
		for res := range in {
			if res.err != nil {
				out <- res
				close(out)
				return
			}
			out <- result{res.value * res.value, nil}
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := generateWithError(2, 3, 0, 5)
	squared := squareWithError((numbers))

	for res := range squared {
		if res.err != nil {
			fmt.Println("Error:", res.err)
			break // NOTE: To prevent showing Result: 0
		}
		fmt.Println("Result:", res.value)
	}
}
