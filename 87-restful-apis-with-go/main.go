package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// NOTE: Building RESTful APIs with Go
/*
	Building RESTful APIs in Go allows you to create web services that communicate over HTTP, providing
	a structured way for clients to interact with your application. This section will cover the essentials
	of designing and implementing a RESTful API.
*/

// NOTE: 1. Understanding REST Principles
/*
	REST (Representational State Transfer) is an architectural style that defines a set of constraints for creating web services. Key principles include:
	•	Statelessness: Each request from a client must contain all the information needed to process it.
	•	Resource-Based: APIs should expose resources (e.g., users, products) and use standard HTTP methods (GET, POST, PUT, DELETE) to interact with them.
*/

// NOTE: 2. Setting Up a RESTful API

/*
	Let’s create a simple RESTful API for managing a collection of books.
*/

// NOTE: 3. Explanation of Key Components
/*
	•	Book Struct: Represents a book resource with fields for ID, Title, and Author. The json tags are used for JSON serialization.
	•	Concurrency Control: A sync.Mutex is used to ensure that access to the books slice is thread-safe when multiple requests are
		processed concurrently.
	•	Handlers:
		•	getBooks: Returns a JSON array of all books.
		•	addBook: Accepts a JSON object in the request body to add a new book.
*/

// NOTE: 4. Testing the API
/*
	You can use tools like Postman or curl to test your API endpoints.
	•	Get All Books:
	```bash
	curl -X GET http://localhost:8080/books
	```

	•	Add a Book:
	```bash
	curl -X POST http://localhost:8080/books/add -d '{"id": 1, "title": "Go Programming", "author": "John Doe"}' -H "Content-Type: application/json"
	```
*/

// NOTE: 5. Additional RESTful Practices
/*
	•	Status Codes: Always return appropriate HTTP status codes based on the outcome of the request
		(e.g., 201 Created for successful creation, 404 Not Found for missing resources).
	•	Error Handling: Return meaningful error messages and status codes when something goes wrong.
	•	Use Resource URLs: Design your API endpoints to reflect the resources they represent (e.g., /books, /books/{id}).
*/

// NOTE: 6. Structuring Your API
/*
	As your API grows, consider structuring your code into separate files or packages to improve organization and maintainability.
		•	Handlers: Move handlers to a separate file (e.g., handlers.go).
		•	Models: Create a models package for your data structures.
		•	Routes: Set up a routes package to manage endpoint registrations.
*/

// NOTE: Key Takeaways:
/*
	•	REST Principles: Understand the fundamental concepts of REST for designing APIs.
	•	JSON Handling: Use Go’s encoding/json package to encode and decode JSON data.
	•	Concurrency: Manage concurrent access to shared resources safely.
	•	HTTP Methods: Implement GET, POST, and other HTTP methods as appropriate.
	•	Testing: Use tools to test your API endpoints.
*/

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var mu sync.Mutex

func getBooks(res http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if books == nil {
		books = []Book{}
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

func addBook(res http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newBook Book
	if err := json.NewDecoder(req.Body).Decode(&newBook); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, newBook)
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newBook)
}

func main() {
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/books/add", addBook)

	fmt.Println("Server is running http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
