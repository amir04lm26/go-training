package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NOTE: Creating Web Applications with Go
/*
	Go is a popular choice for building web applications due to its performance, simplicity, and powerful concurrency model.
	This section will cover the basics of creating a web server, handling requests, and routing.
*/

// NOTE: 1. Setting Up a Simple Web Server
/*
	To create a basic web server in Go, you can use the net/http package. Here’s how to set up a simple server:

	•	http.HandleFunc: Registers a handler function for the specified pattern (in this case, /).
	•	http.ListenAndServe: Starts the server on the specified port.
*/

// NOTE: 2. Handling Requests and Responses
/*
	In the handler function, you can access the request (r) and write responses (w). Here’s a breakdown:
	•	Request: Represents the HTTP request made by the client. It contains information such as headers, URL, and method.
	•	ResponseWriter: Used to construct and send an HTTP response.
*/

func basicHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("reqPath:", req.URL.Path)
	fmt.Fprintf(res, "Hello, %s!", req.URL.Path[1:])
}

func basicServerExample() {
	http.HandleFunc("/", basicHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// NOTE: 3. Routing
/*
	For more complex applications, you may want to use a routing library to manage different routes more easily.
	One popular library is gorilla/mux.
*/

func homeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the home page!")
}

func aboutHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "About us")
}

// NOTE: 4. Handling Query Parameters
/*
	You can access query parameters from the request URL easily:

	You can call this endpoint with a URL like http://localhost:8081/greet?name=Alice.
*/

func greetHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(res, "Hello, %s!", name)
}

func serverWithRoutingExample() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/greet", greetHandler)

	fmt.Println("Server is running on http://localhost:8081")
	http.ListenAndServe(":8081", router)
}

// NOTE: 5. Serving Static Files
/*
	To serve static files like HTML, CSS, and JavaScript, use http.StripPrefix:
	This allows you to access static files in the static directory via URLs starting with /static/.
*/
func serveStaticFilesExample() {
	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("86-simple-server-in-go/static/"))))

	fmt.Println("Server is running on http://localhost:8082")
	http.ListenAndServe(":8082", router)
}

// NOTE: Key Takeaways:
/*
	•	Simple Web Server: Use net/http to create a basic web server.
	•	Request and Response: Handle HTTP requests and responses in your handler functions.
	•	Routing: Use a router like gorilla/mux for managing multiple routes.
	•	Query Parameters: Access query parameters with r.URL.Query().
	•	Static Files: Serve static files using http.FileServer.
*/

func main() {
	go basicServerExample()
	go serverWithRoutingExample()
	serveStaticFilesExample()
}
