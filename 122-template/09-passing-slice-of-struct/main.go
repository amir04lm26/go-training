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
	tpl = template.Must(template.ParseFiles("./09-passing-slice-of-struct/templates/tpl.gohtml"))
}

func main() {
	buddha := sage{"Buddha", "The belief of no beliefs"}
	gandhi := sage{"Gandhi", "be the change"}

	sages := []sage{buddha, gandhi}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
