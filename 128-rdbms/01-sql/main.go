package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// * user:password@tcp(localhost:3306)/db_name?charset=utf8
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/users", users)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "at index")
	check(err)
}

func users(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT fName FROM user_table")
	check(err)

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	// * statement
	stmt, err := db.Prepare("CREATE TABLE customer_table (name VARCHAR(20));")
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer_table", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer_table VALUES ("James");`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	// * return at most one row
	row := db.QueryRow("SELECT name FROM customer_table")
	check(err)

	var name string
	row.Scan(&name)
	fmt.Fprintln(w, "RECEIVED RECORD", name)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer_table SET name="Jimmy" WHERE name="James";`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer_table WHERE name="Jimmy";`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("DROP TABLE customer_table")
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROP TABLE customer_table")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
