package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("./08-passing-struct/templates/tpl.gohtml"))
}

func main() {
	Buddha := sage{"Buddha", "The belief of no beliefs"}

	err := tpl.Execute(os.Stdout, Buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
