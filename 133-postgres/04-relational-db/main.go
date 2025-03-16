package main

// NOTE: create a relation
/*
	```
		CREATE TABLE people(
		    ID              SERIAL PRIMARY KEY       NOT NULL,
		    NAME            TEXT                     NOT NULL,
		    SCORE           INT                      DEFAULT 10 NOT NULL,
		    SALARY          REAL                     NOT NULL
		);

		CREATE TABLE phonenumbers (
			ID          	SERIAL PRIMARY KEY      NOT NULL,
			PHONE       	CHAR(50)                NOT NULL,
			PERSON_ID   	INT                     REFERENCES people(ID) NOT NULL
		);
	```
*/

// NOTE: delete a record
/*
	```
		DELETE FROM phonenumbers WHERE id = 6;
	```
*/

// NOTE: query data and join together
/*
	```
		SELECT p.ID, p.NAME, p.SCORE, p.SALARY, ph.PHONE
		FROM people p
			LEFT JOIN phonenumbers ph ON p.ID = ph.PERSON_ID
		WHERE
			p.ID = 1;
	```
*/

func main() {

}
