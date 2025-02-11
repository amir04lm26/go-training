package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "This is form mcleod")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var dog hotdog
	http.ListenAndServe(":8080", dog)
}
