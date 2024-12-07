package math

import (
	"fmt"
	"testing"
)

// NOTE: Explanation
/*
	•	testing.T: This struct is used to log errors and to signal test failures.
	•	t.Errorf: If the test fails, t.Errorf logs the error and continues running other tests.
	•	Naming: The test function is named TestAdd to indicate it’s testing the Add function.

	•	testing.T: A type provided by the testing package to handle test cases.
		You use it to log errors, and it provides methods like t.Error() and t.Fatalf() for
		reporting failures.
	•	t.Errorf(): Logs a formatted error message but allows the test to continue.
		You can also use t.Fail() or t.FailNow().
*/

func TestAdd(test *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		test.Errorf("Add(2,3) = %d; wants %d", result, expected)
	}
}

// NOTE: Table-Driven Tests
/*
	A common pattern in Go is to write table-driven tests. In this approach, you define a table
	of test cases, each with input and expected output, and iterate over the table to run the tests.
*/

// NOTE: Explanation
/*
	•	Test Table: A slice of structs, where each struct represents a test case with input values
		and the expected result.
	•	Loop: You loop through each test case, run the test, and check the result against the
		expected value.
*/

func TestAddTableDriven(test *testing.T) {
	tests := []struct {
		num1, num2 int
		expected   int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{2, 3, 5},
		{10, 20, 30},
	}

	for _, testCase := range tests {
		result := Add(testCase.num1, testCase.num2)
		if result != testCase.expected {
			test.Errorf("Add(%d, %d) = %d. want %d", testCase.num1, testCase.num2, result, testCase.expected)
		}
	}
}

// NOTE: Benchmarking in Go
/*
	Go allows you to measure the performance of your code using benchmarks. Benchmark functions are
	named BenchmarkXxx and take a pointer to testing.B. They run in a loop to measure the
	performance of the function.

	To run benchmarks, use:
	go test -bench=.

	This will run all benchmarks in the package. You can also filter by name:
	go test -bench=Add

	if you want to display allocation metrics for all benchmarks, use the following command:
	go test -bench=. -benchmem
*/

// NOTE: Explanation
/*
	•	testing.B: This struct allows you to benchmark the code by running it multiple times
		(controlled by b.N, the number of iterations).
	•	b.N: This value is automatically adjusted by the Go testing framework to run the benchmark
		multiple times and provide an accurate measurement.
*/

func BenchmarkAdd(benchmarker *testing.B) {
	for iterator := 0; iterator < benchmarker.N; iterator++ {
		Add(1, 2)
	}
}

// NOTE: Best Practices
/*
	•	Avoid I/O operations: If benchmarking code involves I/O (e.g., file reads/writes or network
		calls), isolate the core logic to avoid skewed results.
	•	Use b.ReportAllocs(): If memory allocations are relevant to your benchmark, call
		b.ReportAllocs() to track them.
	•	Use b.StopTimer() / b.StartTimer(): When setting up heavy data structures, use these to
		exclude the setup time from the benchmark.

	func BenchmarkSetupHeavy(b *testing.B) {
	# Setup
	data := make([]int, 1000)

	b.ResetTimer() // Start timer for benchmark

	for i := 0; i < b.N; i++ {
		_ = len(data)
	}
*/

// NOTE: Test Coverage
/*
	Go provides built-in support for checking test coverage. You can see what percentage of your code
	is covered by tests.

	To check test coverage:
	go test -cover

	You can also generate detailed coverage reports:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

	This will open an HTML file in your browser, showing which lines of code are covered by your tests.
*/

// NOTE: Handling Edge Cases in Tests
/*
	When writing tests, it’s important to cover edge cases (e.g., dividing by zero, empty slices, etc.).
	Let’s look at a couple of examples.
*/

// NOTE: Explanation
/*
	•	Edge Cases: Test cases include dividing by zero, and the test checks whether the error is
		handled properly.
	•	hasError Flag: Indicates whether the test case should result in an error.
*/

func TestDivide(test *testing.T) {
	tests := []struct {
		num1, num2 float64
		expected   float64
		hasError   bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true},
	}

	for _, testCase := range tests {
		result, err := Divide(testCase.num1, testCase.num2)
		if testCase.hasError && err == nil {
			test.Errorf("expected error but got none")
		}
		if !testCase.hasError && result != testCase.expected {
			test.Errorf("Divide(%f, %f) = %f; want %f", testCase.num1, testCase.num2, result, testCase.expected)
		}
	}
}

// NOTE: Sub-tests
/*
	Go supports sub-tests, which allow you to group related tests and run them together.
	This is useful when you have variations of a single test.
*/

// NOTE: Explanation
/*
	•	t.Run: Creates subtests that run individually, making it easy to see which specific case failed.
	•	Test Names: Each subtest has a name, which is shown in the output if the test fails.
*/

func TestAddSubTests(test *testing.T) {
	tests := []struct {
		name       string
		num1, num2 int
		expected   int
	}{
		{"Add positive numbers", 1, 2, 3},
		{"Add zero", 0, 0, 0},
		{"Add negative numbers", -1, -2, -3},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			result := Add(testCase.num1, testCase.num2)
			if result != testCase.expected {
				test.Errorf("Add(%d, %d) = %d; want %d", testCase.num1, testCase.num2, result, testCase.expected)
			}
		})
	}
}

// NOTE: Example Tests
/*
	Go also provides a way to write example tests, which can serve as both documentation and as
	functional tests. These functions are named ExampleXxx.
*/

// NOTE: Explanation
/*
	•	// Output:: This comment specifies the expected output.
		If the actual output doesn’t match, the test fails.
*/
func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output: 3
}
