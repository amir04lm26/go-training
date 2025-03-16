package main

// NOTE: Create tables
/*
	```sql
	CREATE TABLE customers (
	    ID          SERIAL PRIMARY KEY          NOT NULL,
	    FIRST       CHAR(50)                    NOT NULL
	);

	CREATE TABLE movies (
	    ID              SERIAL PRIMARY KEY          NOT NULL,
	    MOVIE           CHAR(50)                    NOT NULL
	);

	CREATE TABLE rentals (
	    ID               SERIAL PRIMARY KEY         NOT NULL,
	    C_ID             INT                        REFERENCES customers (ID) NOT NULL,
	    M_ID             INT                        REFERENCES movies (ID) NOT NULL
	);
	```
*/

// NOTE: INSERT into tables
/*
	```sql
	INSERT INTO customers (FIRST) VALUES ('Amir'), ('Faeze'), ('Bond'), ('Money Penny'), ('X');
	```

	```sql
	INSERT INTO movies (movie) VALUES ('Skyfall'), ('start wars'), ('serial killer'), ('jaws'), ('alien'), ('aliens');
	```

	```sql
	INSERT INTO rentals(c_id, m_id) VALUES (1,4), (5,2), (2,5), (4,1), (3,1), (3,5), (2,1), (4,2);
	```
*/

// NOTE: Query rentals
/*
	```sql
	SELECT * FROM customers;
	SELECT * FROM rentals;
	SELECT * FROM movies;
	```

	```sql
	SELECT c.FIRST, m.MOVIE FROM customers c INNER JOIN rentals r on c.ID = r.C_ID INNER JOIN movies m ON r.M_ID = m.ID;
	```
*/

func main() {

}
