package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/lib/pq"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB
var tpl *template.Template

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

	tpl = template.Must(template.ParseGlob("./14-web-app-insert/template/*.gohtml"))
}

func main() {
	defer db.Close()
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	http.HandleFunc("/books/delete/process", booksDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
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

	// for _, book := range books {
	// 	fmt.Fprintf(w, "%s, %s, %s, $%.2f\r\n", book.Isbn, book.Title, book.Author, book.Price)
	// }

	tpl.ExecuteTemplate(w, "books.gohtml", books)
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

	// fmt.Fprintf(w, "%s, %s, %s, $%.2f\r\n", book.Isbn, book.Title, book.Author, book.Price)
	tpl.ExecuteTemplate(w, "show.gohtml", book)
}

func booksCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

// curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
func booksCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate inputs
	if book.Isbn == "" || book.Title == "" || book.Author == "" || p == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// convert form values
	if f64, err := strconv.ParseFloat(p, 32); err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable)+". Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	} else {
		book.Price = float32(f64)
	}

	// insert values - postgresql place holders
	_, err := db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4);", book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "created.gohtml", book)
}

func booksUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+". You should provide the isbn key", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	// var book Book
	book := Book{} // * The same
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

	tpl.ExecuteTemplate(w, "update.gohtml", book)
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	p := r.FormValue("price")

	if book.Isbn == "" || book.Title == "" || book.Author == "" || p == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if f64, err := strconv.ParseFloat(p, 32); err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable)+". Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	} else {
		book.Price = float32(f64)
	}

	_, err := db.Exec("UPDATE books SET ISBN = $1, TITLE=$2, AUTHOR=$3, PRICE=$4 WHERE isbn = $1",
		book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "updated.gohtml", book)
}

func booksDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+". You should provide the isbn key", http.StatusBadRequest)
		return
	}

	if res, err := db.Exec("DELETE FROM books WHERE isbn = $1", isbn); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		ra, err := res.RowsAffected()

		switch {
		case err != nil:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		case ra == 0:
			http.NotFound(w, r)
			return
		}
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
