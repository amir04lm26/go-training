package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utm-8")

	// * method="get"
	// io.WriteString(w, `
	// 	<form method="GET">
	// 		<input type="text" name="q">
	// 		<input type="submit">
	// 	</form>
	// 	<br>
	// `+v)

	// * method post
	io.WriteString(w, `
		<form method="POST">
			<input type="text" name="q">
			<input type="submit">
		</form>
		<br>
	`+v)

}

func bar(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	s := r.FormValue("subscribe") == "on" // * Check if checked
	w.Header().Set("Content-Type", "text/html; charset=utm-8")
	// * with subscribe
	fmt.Fprintf(w, `
		<form method="POST">
			<span>Name:</span> <input type="text" name="name">
			<br>
			<span>Subscribe:</span> <input type="checkbox" name="subscribe">
			<br>
			<input type="submit">
		</form>
		<br>
		<h2>My name is: %v</h2>
		<h2>Subscribe: %v</h2>
	`, n, s)
}
