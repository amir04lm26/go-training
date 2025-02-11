package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var dog hotdog

	http.ListenAndServe(":8080", dog)
}
