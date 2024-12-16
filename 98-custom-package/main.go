package main

import (
	"fmt"
	"strings"

	mpackage "github.com/amir04lm26/go-custom-pkg"
)

// NOTE: Get the custom package created
/*
	go get github.com/amir04lm26/go-custom-pkg
	go get github.com/amir04lm26/go-custom-pkg@latest
	go get github.com/amir04lm26/go-custom-pkg@b614fff0da3ef337762fdb6bb3ce948cd02d89ef
	go get github.com/amir04lm26/go-custom-pkg@custom-tag
*/

func main() {
	fmt.Println(mpackage.Greet("Amir"))
	fmt.Println(mpackage.Greet(strings.ToUpper("faeze")))
}
