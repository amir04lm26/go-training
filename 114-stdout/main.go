package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, Go")
	fmt.Fprintln(os.Stdout, "Hello, Go")
	io.WriteString(os.Stdout, "Hello, Go\n")
}
