package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotCat int

func (hotCat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	var dog hotdog
	var cat hotCat

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog) // * Applies trailing slash. Also, catch other sub-paths
	mux.Handle("/cat", cat)

	http.ListenAndServe(":8080", mux)
}
