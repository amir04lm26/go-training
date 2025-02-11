package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

// NOTE Handler interface
func (hotdog hotdog) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("From", request.Form, request.PostForm)
	// tpl.ExecuteTemplate(response, "index.gohtml", r.PostForm) // * Only value that come from the post data not query strings
	tpl.ExecuteTemplate(response, "index.gohtml", request.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./01-retrieving-form-data/index.gohtml"))
}

func main() {
	var dog hotdog
	http.ListenAndServe(":8080", dog)
}
