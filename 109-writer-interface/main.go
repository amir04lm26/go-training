package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (person person) writeOut(writer io.Writer) error {
	_, err := writer.Write([]byte(person.name))
	return err
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func main() {
	p1 := person{name: "Amir"}

	file, err := os.Create("./109-writer-interface/output.txt")
	fatalIfErr(err)
	defer file.Close()

	var buffer bytes.Buffer

	err = p1.writeOut(file)
	fatalIfErr(err)

	err = p1.writeOut(&buffer)
	fatalIfErr(err)

	fmt.Println(buffer)
	fmt.Println(buffer.String())
}
