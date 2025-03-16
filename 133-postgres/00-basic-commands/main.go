package main

// NOTE: connect to postgres shell
/*
	```
	psql -d <database name>
	psql -U <username> -d <database name>
	```
*/

// NOTE: quit postgres shell
/*
	```
	\q
	```
*/

// NOTE: Create a Database
/*
	```
	CREATE DATABASE <database name>;
	```
*/

// NOTE: List all databases
/*
	```
	\l
	```
*/

// NOTE: Connect to a database
/*
	```
	\c <database name>
	```
*/

// NOTE: Set Ownership or Permissions
/*
	```
	ALTER DATABASE <database name> OWNER TO <new owner>;
	```
*/

// NOTE: clear logs
/*
	```
	\! clear
	```
*/

// NOTE: Switch back to postgres database
/*
	```
	\c postgres
	```
*/

// NOTE: See current user
/*
	```
	SELECT current_user;
	```
*/

// NOTE: See current database
/*
	```
	SELECT current_database();
	```
*/

// NOTE: drop (remove, delete) database
/*
	```
	DROP DATABASE <database name>;
	```
*/

// NOTE: Users detail
/*
	```
	\du
	```
*/

func main() {

}
