package main

import "fmt"

// NOTE: Pipeline Processing
/*
	Pipelines allow you to process data in multiple stages, where the output of one stage becomes
	the input for the next stage. In Go, pipelines can be implemented using goroutines and channels.
*/

// NOTE: Explanation
/*
	•	generate: Generates numbers and sends them through the channel.
	•	square: Receives numbers from the generate channel, squares them, and sends the result to
		the next stage.
	•	double: Receives squared numbers, doubles them, and sends the result to the next stage.
	•	Pipeline: The numbers flow through the pipeline, where each stage applies some transformation.
*/

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := generate(2, 3, 4, 5)
	squared := square(numbers)
	doubled := double(squared)

	for result := range doubled {
		fmt.Println(result)
	}
}
