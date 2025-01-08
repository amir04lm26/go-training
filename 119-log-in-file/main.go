package main

import (
	"fmt"
	"log"
	"os"
)

const currentDirectory = "./119-log-in-file"

func main() {
	file, err := os.Create(currentDirectory + "/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	secondFile, err := os.Open(currentDirectory + "/not-exist.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer secondFile.Close()

	fmt.Printf("Check the %s/log.txt file in the directory\n", currentDirectory)
}
