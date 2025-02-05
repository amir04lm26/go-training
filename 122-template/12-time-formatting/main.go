package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

var funcMap = template.FuncMap{
	"fDateMDY": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("./12-time-formatting/templates/tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
