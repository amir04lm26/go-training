package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB

// NOTE: SQL Injection prevention - sql package
/*
	Behind the scenes, db.QueryRow (and also db.Query() and db.Exec()) work by creating a new prepared statement on the database,
	and subsequently execute that prepared statement using the placeholder parameters provided. This means that all three methods
	are safe from SQL injection when used correctly

	Syntax:
	```sql
	PREPARE name [ ( data_type [, ...] ) ] AS statement
	```


	Example:
	```sql
	PREPARE fooplan (int, text, bool, numeric) AS
    INSERT INTO foo VALUES($1, $2, $3, $4);
	EXECUTE fooplan(1, 'Hunter Valley', 't', 200.00);
	```
*/

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:123456@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection Stablished")
}

func main() {
	defer db.Close()
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.ListenAndServe(":8080", nil)
}

// http://localhost:8080/books
func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price) // * Order matters
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\r\n", book.Isbn, book.Title, book.Author, book.Price)
	}
}

// http://localhost:8080/books/show?isbn=978-1505255607
func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	book := Book{}
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	switch {
	case err == sql.ErrNoRows:
		// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\r\n", book.Isbn, book.Title, book.Author, book.Price)
}
