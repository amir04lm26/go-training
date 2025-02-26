package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// NOTE: unmarshal vs decode
/*
	In Go’s encoding/json package, both json.Unmarshal and json.Decoder are used for parsing JSON,
	but they serve different purposes and have different use cases:
*/

// NOTE: 1. json.Unmarshal
/*
	•	Use case: When you have a complete JSON document in memory (as a []byte) and want to convert
		it into a Go struct or other data type.

	•	Method signature:
		```go
			func Unmarshal(data []byte, v interface{}) error
		```

	•	Pros:
		•	Simple and easy to use when dealing with small JSON payloads.
		•	Works well when JSON is already loaded into memory.
	•	Cons:
		•	Requires the entire JSON to be loaded into memory.
		•	Not efficient for streaming large JSON payloads.
*/

// NOTE: 2. json.Decoder (Decode method)
/*
	•	Use case: When working with JSON streams (e.g., reading from a file or network), allowing
		incremental parsing.

	•	Method signature:
		```go
		func (dec *Decoder) Decode(v interface{}) error
		```

	•	Pros:
		•	Works well for large JSON data that is streamed or read from a file.
		•	Efficient because it processes JSON incrementally, reducing memory overhead.
	•	Cons:
		•	Can be slightly more complex than Unmarshal for simple use cases.
		•	Cannot easily handle JSON fragments (e.g., reading part of a JSON array).
*/

// NOTE: When to Use Which?
/*
	Use Case								Preferred Method
	Small JSON payloads in memory			json.Unmarshal
	Large JSON data or streaming input		json.Decoder.Decode

	Use json.Unmarshal for simple cases where the entire JSON is available in memory.
	Use json.Decoder when reading JSON from a file, network, or other streaming sources to save
	memory and improve efficiency.
*/

type person struct {
	FName string
	LName string
	Items []string
}

type imageThumb struct {
	URL    string `json:"Url"`
	Height int    `json:"Height"`
	Width  int    `json:"Width"`
}

type image struct {
	// Width     int        `json:"Width"`
	// Height    int        `json:"Height"`
	// Title     string     `json:"Title"`
	Thumbnail imageThumb `json:"Thumbnail"`
	// Animated  bool       `json:"Animated"`
	IDs []int `json:"IDs,omitempty"`
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="utf-8">
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<title>Foo</title>
			</head>
			<body>
				<h1>Foo</h1>
			</body>
			</html>
	`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(bs)
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	var data image
	rData := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rData), &data)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(data)

	for index, id := range data.IDs {
		fmt.Println(index, id)
	}

	fmt.Println(data.Thumbnail.URL)
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	var data image
	rData := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.NewDecoder(strings.NewReader(rData)).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(data)

	for index, id := range data.IDs {
		fmt.Println(index, id)
	}

	fmt.Println(data.Thumbnail.URL)
}
