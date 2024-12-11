package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // * Import the pq driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NOTE: Working with Databases in Go
/*
	Interacting with databases is a common requirement for web applications.
	In Go, you can use the database/sql package along with specific drivers for different databases.
	This section will cover the basics of connecting to a database, executing queries, and handling results.
*/

// NOTE: 1. Setting Up a Database Connection
/*
	To work with a database, you first need to install the appropriate driver for your database. For example,
	to work with PostgreSQL, you would use pgx or lib/pq.

	Installation for PostgreSQL:
	```bash
	go get -u github.com/lib/pq
	```
*/

func connectToDataBase() *sql.DB {
	connStr := "postgres://amir:123456@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// * Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	return db
}

// NOTE: 2. Executing Queries
/*
	Once connected, you can execute SQL queries to manipulate data.
*/

func createBooksTableIfNotExists(db *sql.DB) {
	checkTableQuery := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_name = 'books'
		);
	`

	var exists bool
	err := db.QueryRow(checkTableQuery).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking if table exits: %v", err)
	}

	if !exists {
		createTableQuery := `
			CREATE TABLE books (
				id SERIAL PRIMARY KEY,
				title TEXT NOT NULL,
				author TEXT NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
		`
		_, err := db.Exec(createTableQuery)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
		fmt.Println("Table 'books' created successfully!")
	} else {
		fmt.Println("Table 'books' already exists.")
	}
}

// NOTE: Inserting Data:
func insertBook(db *sql.DB, title, author string) {
	// query := "INSERT INTO books (title, author) VALUES ($1, $2)"
	// _, err := db.Exec(query, title, author)
	stmt, err := db.Prepare("INSERT INTO books (title, author) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, author)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book inserted successfully!")
}

// NOTE: 3. Querying Data
/*
	To retrieve data, use the Query method.
*/

func getBooks(db *sql.DB) {
	rows, err := db.Query("SELECT title, author FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title, author string
		if err := rows.Scan(&title, &author); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Title: %s, Author: %s\n", title, author)
	}
}

// NOTE: 4. Handling Errors
/*
	Always check for errors after executing database operations to handle any issues gracefully.
*/

// NOTE: 5. Structuring Queries
/*
	You can use prepared statements for more complex queries or to avoid SQL injection:

	Prepared Statement Example:
	```bash
	stmt, err := db.Prepare("INSERT INTO books (title, author) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Go Programming", "John Doe")
	```
*/

// NOTE: 6. Using ORM Libraries
/*
	For larger projects, you might consider using an Object-Relational Mapping (ORM) library like GORM or Ent.
	These libraries provide higher-level abstractions to work with databases.

	Installation:
	```bash
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	```
*/

// NOTE: Key Takeaways:
/*
	•	Database Connection: Use the database/sql package and the appropriate driver to connect to your database.
	•	Executing Queries: Use Exec for inserts and Query for retrieving data.
	•	Prepared Statements: Use prepared statements to avoid SQL injection.
	•	ORM Libraries: Consider using an ORM for easier database interactions.
*/

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ormExample() {
	connStr := "postgres://amir:123456@localhost:5432/testdb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("Connected to database successfully")

	// * Migrate the schema
	db.AutoMigrate(&Book{})

	// * Create
	// db.Create(&Book{Title: "New GO Programming", Author: "John Doe"})

	// * Read
	var book Book
	db.First(&book, 1) // * find book with ID 1

	fmt.Println("First Book:", book)

	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Fatalf("Error reading books: %v", result.Error)
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func main() {
	// db := connectToDataBase()
	// defer db.Close()

	// createBooksTableIfNotExists(db)

	// // insertBook(db, "Go Programming", "John Doe")

	// getBooks(db)

	ormExample()
}
