package tpl

import (
	"log"
	"path/filepath"
	"text/template"
)

var Tpl *template.Template

func LoadTemplates() {
	var err error
	Tpl, err = template.ParseGlob(filepath.Join("web", "template", "*.gohtml"))
	if err != nil {
		log.Fatalf("Error parsing templates: %v\r\n", err)
	}
}
