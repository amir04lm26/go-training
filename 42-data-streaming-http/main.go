package main

import (
	"fmt"
	"io"
	"net/http"
)

// NOTE: Streaming HTTP Responses
/*
	When dealing with HTTP APIs, sometimes you’ll want to stream the response incrementally, especially when
	the response is too large to fit in memory.
*/

// NOTE: Explanation
/*
	•	http.Get: Initiates an HTTP GET request.
	•	resp.Body.Read: Reads the HTTP response body in chunks, allowing the response to be processed without
		loading the entire content into memory.
*/

func main() {
	resp, err := http.Get("https://apix.snappshop.ir/landing/v1/megamenu?lat=35.7435756&lng=51.3306158")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	buffer := make([]byte, 1024) // READ 1KB at a time
	for {
		numberOfBytes, err := resp.Body.Read(buffer)
		if numberOfBytes > 0 {
			fmt.Println(string(buffer[:numberOfBytes]))
		}
		if err == io.EOF {
			break // End of stream
		}
	}
}
