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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// type items struct {
// 	Wisdom    []sage
// 	Transport []car
// }

func init() {
	tpl = template.Must(template.ParseFiles("./10-passing-struct-slice-of-struct/templates/tpl.gohtml"))
}

func main() {
	buddha := sage{"Buddha", "The belief of no beli efs"}
	gandhi := sage{"Gandhi", "be the change"}
	sages := []sage{buddha, gandhi}

	ford := car{"Ford", "F150", 2}
	toyota := car{"Toyota", "Corolla", 4}
	cars := []car{ford, toyota}

	// items := items{Wisdom: sages, Transport: cars}
	// items := items{sages, cars}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
