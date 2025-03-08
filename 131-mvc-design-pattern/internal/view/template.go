package view

import (
	"html/template"
	"log"

	// "path" // * for handling URLs or predefined paths in a cross-platform way
	"path/filepath" // * for filesystem paths
)

var Tpl *template.Template

func LoadTemplates() {
	var err error
	Tpl, err = template.ParseGlob(filepath.Join("internal", "view", "*.gohtml"))
	if err != nil {
		log.Fatalf("Error parsing templates: %v\r\n", err)
	}
}
