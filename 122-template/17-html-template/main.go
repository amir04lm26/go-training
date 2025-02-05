package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Page struct {
	Title   string
	Heading string
	Input   string
}

func init() {
	tpl = template.Must(template.ParseGlob("./17-html-template/templates/*.gohtml"))
}

func main() {
	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!")</script>`,
	}

	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
