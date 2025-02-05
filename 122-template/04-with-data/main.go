package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./04-with-data/templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Amir Zare")
	if err != nil {
		log.Fatalln(err)
	}
}
