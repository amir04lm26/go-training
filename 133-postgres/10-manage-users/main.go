package main

// NOTE: users & privileges
/*
	see current user

	```sql
	SELECT current_user;
	```
*/

// NOTE: details of users
/*
	```sql
	\du
	```
*/

// NOTE: create user
/*
	```sql
	CREATE USER james WITH PASSWORD 'password';
	```
*/

// NOTE: grant privileges
/*
	privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL

	```
	GRANT ALL PRIVILEGES ON DATABASE company to james;
	```
*/

// NOTE: revoke privileges
/*
	```sql
	REVOKE ALL PRIVILEGES ON DATABASE company from james;
	```
*/

// NOTE: alter
/*
	```sql
	ALTER USER james WITH SUPERUSER;
	```

	```sql
	ALTER USER james WITH NOSUPERUSER;
	```
*/

// NOTE: remove
/*
	```sql
	DROP USER james;
	```
*/

// NOTE: Comparison: GRANT ALL PRIVILEGES vs SUPERUSER
/*
	Feature					GRANT ALL PRIVILEGES								SUPERUSER
	Scope					Specific database object (e.g., table, schema).		Entire database cluster.
	Privileges				All privileges on the specified object.				All privileges on all objects + administrative privileges.
	Administrative Access	No.	Yes.
	Use Case				Grant full access to a specific object.				Grant full administrative control.
	Security Risk			Limited to the specified object.					High (unrestricted access).
*/

func main() {

}
