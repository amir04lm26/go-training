package main

import (
	"errors"
	"fmt"
	"log"
)

// NOTE: Error Handling in Go
/*
	In Go, error handling is done explicitly using the error type.
	Functions that can fail return an error as their last return value, and it’s the
	responsibility of the caller to handle the error.
*/

func greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greeting("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
