package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NOTE: RESTful Services in Go
/*
	You can build RESTful services using Go’s net/http package. Here’s a simple example:
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser(resWriter http.ResponseWriter, _ *http.Request) {
	user := User{ID: 1, Name: "Alice"}
	resWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resWriter).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
