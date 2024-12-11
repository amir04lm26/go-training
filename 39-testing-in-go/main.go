package main

// NOTE: Testing in Go
/*
	Go has a built-in testing package that provides support for automated testing.
	You can write unit tests, benchmarks, and example tests to ensure your program works as expected.
*/

// NOTE: Writing a Basic Unit Test
/*
	A unit test checks the behavior of a small unit of your program, usually a single function.
	Test functions in Go are named TestXxx (where Xxx is descriptive), and they take a pointer to
	testing.T.
*/

// NOTE: Running Tests
/*
	To run tests, use the go test command in the directory where the test files are located.

	go test

	This will automatically find all *_test.go files and run their tests.

	To see detailed output, you can add the -v flag:
	```bash
	go test -v
	```
*/
