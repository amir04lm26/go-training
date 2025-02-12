package main

import (
	"io"
	"net/http"
)

func handleDog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func handleCat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func handleLion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "lion lion lion")
}

func main() {

	// * Adds to t he DefaultServeMux
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/cat", handleCat)

	http.HandleFunc("/mouse", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "mouse mouse mouse")
	})
	http.Handle("/lion", http.HandlerFunc(handleLion))

	http.ListenAndServe(":8080", nil) // * nil -> uses the DefaultServeMux
}
