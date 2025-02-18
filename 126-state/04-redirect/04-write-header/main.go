package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at foo ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at bar ", r.Method, "\n\n")
	// process form data
	w.Header().Add("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at barred ", r.Method, "\n\n")
	tpl.Execute(w, "index.gohtml")
}
