package main

// NOTE: Data hierarchy
/*
	RDBS/RDBMS -> database -> table -> records -> fields -> characters
*/

// NOTE: See the database details
/*
	```
	\d
	```
*/

// NOTE: Create a Table
/*
	```
	CREATE TABLE employees (
		ID 			   INT PRIMARY KEY     	NOT NULL,
		NAME           TEXT    				NOT NULL,
		RANK           INT     				NOT NULL,
		ADDRESS        CHAR(50),
		SALARY         REAL 				DEFAULT 25500.00,
		BDAY		   DATE 				DEFAULT '1900-01-01'
	);
	```
*/

// NOTE: See table details
/*
	```
	\d <table name>
	```
*/

// NOTE: drop a table
/*
	```
	DROP TABLE <table name>;
	```
*/

func main() {

}
