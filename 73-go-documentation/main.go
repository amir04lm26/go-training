// Package main talks about documentation
package main

import "fmt"

// NOTE: Go Documentation: go doc and godoc
/*
	Go provides built-in tools for generating and viewing documentation:
		•	go doc: Command-line tool to view documentation for Go code.
		•	godoc: A web-based server that serves documentation for Go projects.

	Example: Using go doc
	```bash
	go doc fmt.Println
	```

	Show the source code
	```bash
	go doc -src fmt.Println
	```

	This command shows the documentation for the fmt.Println function.

	Example: Generating Web Docs with godoc
	```bash
	godoc -http=:6060
	```

	Install using:
	```bash
	go install golang.org/x/tools/cmd/godoc@latest
	```

	This will launch a local web server on localhost:6060 where you can browse your
	project’s documentation.
*/

// Sum adds inputs together
// the rest of description goes here
func Sum(num ...int) int {
	var total int
	for _, singleNum := range num {
		total += singleNum
	}
	return total
}

func main() {
	fmt.Println(Sum(2, 23, 3))
}
