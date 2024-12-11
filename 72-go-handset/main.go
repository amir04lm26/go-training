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

// NOTE: Updating Dependencies
/*
	To update your dependencies to the latest version, you can use:
	```bash
	go get -u
	```

	To update a specific dependency, use:
	```bash
	go get example.com/package@latest
	```
*/

// NOTE: Versioning
/*
	Go modules support semantic versioning. You can specify a version when adding a dependency:
	```bash
	go get example.com/package@v1.2.3
	```

	You can also use version ranges, such as:
	```bash
	go get example.com/package@^1.0.0
	```
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

// NOTE: Key Takeaways:
/*
	•	Module Initialization: Use go mod init to create a new module.
	•	Automatic Dependency Management: Go automatically updates go.mod when you import new packages.
	•	Updating and Removing Dependencies: Use go get -u to update and go mod tidy to clean up.
	•	Versioning Support: Go modules support semantic versioning.
*/

// NOTE: Running a Go Application
/*
	You can run your Go application directly using the go run command. This command compiles and executes the Go files specified.

	Example:
	```bash
	go run main.go
	```

	You can also run multiple files at once:
	```bash
	go run main.go helper.go
	```
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

// NOTE: Installing a Go Package
/*
	If you want to install a Go package globally (making it available for all projects), you can use the go install command. This compiles the package and places the executable in your GOPATH/bin directory.

	Example:
	```bash
	go install example.com/myproject@latest
	```
*/

// NOTE: Key Takeaways:
/*
	•	Building Applications: Use go build to compile your Go code into an executable.
	•	Running Applications: Use go run to execute your Go files without creating an executable.
	•	Cross-Compiling: Easily compile your application for different platforms using GOOS and GOARCH.
	•	Installing Packages: Use go install to make your package globally available.
	•	Running Tests: Execute tests using the go test command.
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
