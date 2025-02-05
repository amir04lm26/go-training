package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseGlob("./13-predefined-global-functions/templates/*.gohtml"))
}

func main() {
	numeric := []string{"zero", "one", "two", "three", "four", "five"}

	user1 := user{"Buddha", "The belief of no beliefs", false}
	user2 := user{"Gandhi", "Be the change", true}
	user3 := user{"", "Nobody", true}

	data := struct {
		Numeric []string
		Users   []user
	}{
		numeric, []user{user1, user2, user3},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
