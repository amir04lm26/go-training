package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.png", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.png">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("./asset/toby.png")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return // * Prevent moving forward and crashing
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
}
