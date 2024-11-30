package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

// NOTE
/*
	go get <package-path> // install a package
	go get -u <package-path> // update a package
	go mod tidy // to clean up any unused dependencies from the go.mod and go.sum files
*/

func main() {
	fmt.Println(quote.Go())
}
