package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// * Both works
	// tpl = template.Must(template.ParseGlob("./05-assigning-variables/templates/*"))
	tpl = template.Must(template.ParseFiles("./05-assigning-variables/templates/tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Amir JZ")
	if err != nil {
		log.Fatalln(err)
	}
}
