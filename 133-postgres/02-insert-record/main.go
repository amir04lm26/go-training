package main

// NOTE: Insert a record
/*
	```
	INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
	```
*/

// NOTE: List records in a table
/*
	```sql
	SELECT * FROM <table name>;
	```

	Example:
	```sql
		SELECT people.NAME FROM people;

		or with alias

		SELECT p.NAME FROM people p;
	```
*/

// NOTE: Insert a record - variations
/*
	omitted values will have the default value:
	```
	INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
	```

	we can use DEFAULT rather leaving a field blank or specifying a value:
	INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
*/

// NOTE: Insert multiple rows
/*
	```
	INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
	```
*/

func main() {

}
