package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./15-nested-template/templates/*.gohtml"))
}

func main() {
	data := struct {
		Description  string
		BabyBearName string
	}{
		Description:  "polar bear(nested template)",
		BabyBearName: "Alex",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
