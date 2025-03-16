package main

// NOTE: 1. INNER JOIN (JOIN: Default INNER JOIN)
/*
	The INNER JOIN returns only the rows that have matching values in both tables.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		INNER JOIN table2
		ON table1.column = table2.column;
	```

	Basic Example:
	```sql
		SELECT * FROM people INNER JOIN phonenumbers ON people.ID = phonenumbers.PERSON_ID;

		SELECT people.NAME, phonenumbers.PHONE
		FROM people
			INNER JOIN phonenumbers ON people.ID = phonenumbers.PERSON_ID
	```

	Example (alias):
	```sql
		SELECT p.ID, p.NAME, ph.PHONE
		FROM people p
		INNER JOIN phonenumbers ph ON p.ID = ph.PERSON_ID;
	```

	Behavior:
	 - Returns only rows where there is a match in both tables.
	 - If there is no match, the row is excluded from the result.

	Use Case:
	 - When you only want records that have corresponding data in both tables.
*/

// NOTE: 2. LEFT JOIN (or LEFT OUTER JOIN)
/*
	The LEFT JOIN returns all rows from the left table and the matching rows from the right table. If there is no match, the result will include NULL values for the right table's columns.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		LEFT JOIN table2
		ON table1.column = table2.column;
	```

	Example:
	```sql
		SELECT p.ID, p.NAME, ph.PHONE
		FROM people p
		LEFT JOIN phonenumbers ph ON p.ID = ph.PERSON_ID;

		SELECT people.NAME, phonenumbers.PHONE
		FROM people
			INNER JOIN phonenumbers ON people.ID = phonenumbers.PERSON_ID
	```

	Behavior:
	 - Returns all rows from the left table (people), even if there is no match in the right table (phonenumbers).
	 - If there is no match, the right table's columns will contain NULL.

	Use Case:
	- When you want all records from the left table, regardless of whether they have matching records in the right table.
	- Useful for finding rows in the left table that do not have corresponding rows in the right table.
*/

// NOTE: 3. RIGHT JOIN (or RIGHT OUTER JOIN)
/*
	The RIGHT JOIN returns all rows from the right table and the matching rows from the left table.
	If there is no match, the result will include NULL values for the left table's columns.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		RIGHT JOIN table2
		ON table1.column = table2.column;
	```

	Example:
	```sql
		SELECT p.ID, p.NAME, ph.PHONE
		FROM people p
		RIGHT JOIN phonenumbers ph ON p.ID = ph.PERSON_ID;
	```

	Behavior:
	 - Returns all rows from the right table (phonenumbers), even if there is no match in the left table (people).
	 - If there is no match, the left table's columns will contain NULL.

	Use Case:
	 - When you want all records from the right table, regardless of whether they have matching records in the left table.
	 - Rarely used compared to LEFT JOIN.
*/

// NOTE: 4. FULL OUTER JOIN
/*
	The FULL OUTER JOIN returns all rows when there is a match in either the left or right table.
	If there is no match, the missing side will contain NULL.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		FULL OUTER JOIN table2
		ON table1.column = table2.column;
	```

	Example:
	```sql
		SELECT p.ID, p.NAME, ph.PHONE
		FROM people p
		FULL OUTER JOIN phonenumbers ph ON p.ID = ph.PERSON_ID;
	```

	Behavior:
	 - Returns all rows from both tables.
	 - If there is no match, the missing side will contain NULL.

	Use Case:
	 - When you want all records from both tables, including unmatched rows from either table.
*/

// NOTE: 5. CROSS JOIN
/*
	The CROSS JOIN returns the Cartesian product of the two tables, meaning it combines each
	row of the first table with every row of the second table.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		CROSS JOIN table2;
	```

	Basic Example:
	```sql
		SELECT * FROM employees CROSS JOIN phonenumbers;
	```

	Example:
	```sql
		SELECT p.NAME, ph.PHONE
		FROM people p
		CROSS JOIN phonenumbers ph;
	```

	Behavior:
	 - Returns all possible combinations of rows from both tables.
	 - No ON clause is used because there is no condition to match rows.

	Use Case:
	 - When you need to generate all possible combinations of rows (e.g., for testing or generating a matrix).
*/

// NOTE: 6. SELF JOIN
/*
	A SELF JOIN is a regular join, but the table is joined with itself.
	It is useful for querying hierarchical data or comparing rows within the same table.

	Syntax:
	```sql
		SELECT columns
		FROM table1 t1
		JOIN table1 t2
		ON t1.column = t2.column;
	```

	Example:
	Suppose you have an employees table with a manager_id column that references the ID of another employee.
	To find employees and their managers:
	```sql
		SELECT e1.NAME AS Employee, e2.NAME AS Manager
		FROM employees e1
		LEFT JOIN employees e2 ON e1.manager_id = e2.ID;
	```

	Behavior:
	 - Treats the same table as two separate tables for the purpose of the join.
	 - Can use any type of join (INNER, LEFT, etc.).

	Use Case:
	 - When you need to compare rows within the same table or query hierarchical relationships.
*/

// NOTE: 7. NATURAL JOIN
/*
	The NATURAL JOIN automatically joins tables based on columns with the same name and data type.
	It is not commonly used because it can lead to unexpected results.

	Syntax:
	```sql
		SELECT columns
		FROM table1
		NATURAL JOIN table2;
	```

	Example:
	```sql
		SELECT *
		FROM people
		NATURAL JOIN phonenumbers;
	```

	Behavior:
	 - Automatically matches columns with the same name in both tables.
	 - Does not require an ON clause.

	Use Case:
	 - Rarely used due to its implicit behavior and potential for errors.
*/

// NOTE: Summary of joins
/*
	Join Type			Description																		Use Case
	INNER JOIN			Returns only matching rows from both tables.									When you need only matching records.
	LEFT JOIN			Returns all rows from the left table and matching rows from the right table.	When you want all records from the left table, even if unmatched.
	RIGHT JOIN			Returns all rows from the right table and matching rows from the left table.	When you want all records from the right table, even if unmatched.
	FULL OUTER JOIN		Returns all rows from both tables, with NULL for unmatched rows.				When you want all records from both tables, including unmatched rows.
	CROSS JOIN			Returns the Cartesian product of both tables.									When you need all possible combinations of rows.
	SELF JOIN			Joins a table with itself.														When querying hierarchical data or comparing rows within the same table.
	NATURAL JOIN		Automatically joins tables based on columns with the same name.					Rarely used due to implicit behavior.
*/

func main() {

}
