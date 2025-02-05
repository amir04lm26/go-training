package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// NOTE: The file extension don't matter for parsing
	// tpl, err := template.ParseFiles("./01-parse-files/tpl.gohtml", "./tpl2.gohtml")
	tpl, err := template.ParseFiles("./01-parse-files/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("./01-parse-files/tpl2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("./01-parse-files/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// err = tpl.Execute(os.Stdout, nil)
	err = tpl.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
