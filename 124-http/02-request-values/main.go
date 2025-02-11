package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

// NOTE Handler interface
func (hotdog hotdog) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		request.Method,
		request.URL,
		request.Form,
		request.Header,
		request.Host,
		request.ContentLength,
	}
	tpl.ExecuteTemplate(response, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./02-request-values/index.gohtml"))
}

func main() {
	var dog hotdog
	http.ListenAndServe(":8080", dog)
}
