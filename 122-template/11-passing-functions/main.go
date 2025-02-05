package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

var funcMap = template.FuncMap{
	"tu": strings.ToUpper,
	"ft": firstTree,
}

// NOTE: Only data formatting helper functions should pass to the template in order to respect MVC pattern

func init() {
	// * Why there is a `New` in the library? We need it to set a new for the template when using `Parse`
	tplNew := template.Must(template.New("something").Parse("This is my string template\n"))
	tplNew.ExecuteTemplate(os.Stdout, "something", nil)

	// * We have to first define the functions then we can use them in using ParseFiles method
	// NOTE: name that passes to the `New` method should match with the `ParseFiles` file name as they're not in the same directory
	tpl = template.Must(template.New("tpl.gohtml").Funcs(funcMap).ParseFiles("./11-passing-functions/templates/tpl.gohtml"))
}

func firstTree(str string) string {
	str = strings.TrimSpace(str)
	str = str[:3]
	return str
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
