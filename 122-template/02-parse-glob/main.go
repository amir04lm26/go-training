package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// NOTE: The file extension don't matter for parsing
	// tpl, err := template.ParseFiles("./01-parse-file/tpl.gohtml", "./tpl2.gohtml")
	tpl, err := template.ParseGlob("./02-parse-glob/templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
