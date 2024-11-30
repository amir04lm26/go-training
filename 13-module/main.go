package main

import (
	"fmt"
	"go/go-examples/13-module/greeting"
)

// NOTE: Edit the example.com/hello module to use your local example.com/greetings module.
// go mod edit -replace example.com/greetings=../greetings

func main() {
	fmt.Println(greeting.Hello("Amir"))
}
