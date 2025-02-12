package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/toby.png", toby)
	http.HandleFunc("/dog-pic", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<!--not serving from our server-->
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func cat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<!--image doesn't serve-->
		<img src="/toby.jpg">
	`)
}

// *  Only this one is serving
func toby(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./asset/toby.png")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	defer f.Close()

	io.Copy(w, f)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<img src="/toby.png">
		<img src="toby.png">
	`)
}
