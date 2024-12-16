package main

import (
	"fmt"
	"log"
	"os"
)

// NOTE: The init function¶
/*
	Finally, each source file can define its own niladic init function to set up whatever state is required.
	(Actually each file can have multiple init functions.) And finally means finally: init is called after
	all the variable declarations in the package have evaluated their initializers, and those are evaluated
	only after all the imported packages have been initialized.

	Besides initializations that cannot be expressed as declarations, a common use of init functions is to
	verify or repair correctness of the program state before real execution begins.
*/

// NOTE: niladic: Of an operator or function in a program, having no arguments

var (
	user = os.Getenv("USER")
)

func init() {
	fmt.Println("Begin initialization")
	fmt.Println("user:", user)

	if user != "amirjz" {
		log.Fatal("$USER is not correct")
	}
}

func main() {
	fmt.Println("Begin main logic")
}
