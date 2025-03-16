package main

// NOTE:  where
/*
	Adding **WHERE** to a SQL query allows you to filter results.

	```sql
	SELECT * FROM employees WHERE salary > 60000;
	```
*/

// NOTE:  and
/*
	```sql
	SELECT * FROM employees WHERE salary > 60000 AND score = 26;
	```
*/

// NOTE: in
/*
	```sql
	SELECT * FROM employees WHERE score IN (25,26);
	```
*/

// NOTE: not
/*
	```sql
	SELECT * FROM employees WHERE score NOT IN (25,26);
	```
*/

// NOTE: between
/*
	```sql
	SELECT * FROM employees WHERE score BETWEEN 23 AND 26;
	```
*/

// NOTE: is not null
/*
	```sql
	SELECT * FROM employees WHERE score IS NOT NULL;
	```
*/

// NOTE: like
/*
	% -> anything
	```sql
	SELECT * FROM employees WHERE name LIKE '%an%';
	```
*/

// NOTE: or
/*
	```sql
	SELECT * FROM employees WHERE score <= 24 OR salary < 50000;
	```
*/

// NOTE: limit
/*
	Limit the number of records returned

	```sql
	SELECT * FROM employees LIMIT 4;

	SELECT * FROM employees ORDER BY id LIMIT 4;
	```
*/

func main() {

}
