package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

// NOTE: SQL
/*
	```sql
	CREATE TABLE books (
		ISBN CHAR(14) PRIMARY KEY NOT NULL,
		TITLE VARCHAR(255) NOT NULL,
		AUTHOR VARCHAR(255) NOT NULL,
		PRICE DECIMAL(5, 2) NOT NULL
	);
	```
*/

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection Stablished")

	rows, err := db.Query("SELECT * FROM BOOKS")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.isbn, &book.title, &book.author, &book.price) // * Order matters
		if err != nil {
			// log.Println(err)
			panic(err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Printf("%s, %s, %s, $%.2f\r\n", book.isbn, book.title, book.author, book.price)
	}
}
