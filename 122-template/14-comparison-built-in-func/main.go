package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./14-comparison-built-in-func/templates/*.gohtml"))
}

func main() {
	data := struct {
		Score1 int
		Score2 int
	}{7, 9}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
