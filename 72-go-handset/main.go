package main

// NOTE: Go Modules
/*
	Go modules manage dependencies in Go projects. A module is a collection of Go packages
	stored in a directory with a go.mod file at its root.
*/

// NOTE: Example: Initializing a Go Module
/*
	``` bash
	go mod init example.com/my-module
	```
*/

// NOTE: Adding Dependencies
/*
	You can add external packages by running:
	``` bash
	go get github.com/pkg/errors
	```

	This will add the errors package to your go.mod file.
*/

// NOTE: Example: go.mod File
/*
	module example.com/my-module

	go 1.18

	require github.com/pkg/errors v0.9.1
*/

// NOTE: Explanation:
/*
	•	go.mod: Defines the module and its dependencies.
	•	go get: Adds a dependency to the project.
*/

// NOTE: Building Go Programs
/*
	Go programs are typically built using the go build command.
	This compiles the code into an executable binary.

	```bash
	go build
	```

	This command compiles the Go code in the current directory and produces an executable file
	named after the module (or the directory name).

	Example: Building a Go Program
	```bash
	go build -o myprogram main.go
	```

	This command will produce an executable named myprogram.

	You can cross-compile your Go application for different operating systems and architectures.
	For example, to build for Windows, you can run:

	```bash
	GOOS=windows GOARCH=amd64 go build
	```
*/

// NOTE: Go Tools Overview
/*
	Go comes with a set of powerful tools for managing and developing code.
		•	go fmt: Formats your code according to Go’s style guidelines.
		•	go vet: Examines your code for potential issues.
		•	go test: Runs tests in your project.
		•	go mod tidy: Cleans up your go.mod file by removing unused dependencies and include new dependencies.
		•	go run: Runs a Go program without building an executable.
*/

func main() {

}
