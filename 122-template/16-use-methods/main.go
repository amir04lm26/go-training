package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArgs(input int) int {
	return input * 2
}

func (p *person) helper() string {
	return strings.ToUpper(p.Name)
}

func init() {
	tpl = template.Must(template.ParseGlob("./16-use-methods/templates/*.gohtml"))
}

func main() {
	data := person{"Amir", 27}
	fmt.Println(data.helper())

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
