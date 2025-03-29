package main

// NOTE: CAP Theorem (Consistency, Availability, Partition Tolerance)
/*
	The CAP Theorem, proposed by Eric Brewer, states that in a distributed database,
	you can only achieve two out of three guarantees at any given time:

	Consistency (C)
	Every read receives the most recent write or an error.
	Note that consistency as defined in the CAP theorem is quite different from the
	consistency guaranteed in ACID database transactions.

	Availability (A)
	Every request received by a non-failing node in the system must result in a response.
	This is the definition of availability in CAP theorem as defined by Gilbert and Lynch.
	Note that availability as defined in CAP theorem is different from high availability
	in software architecture.

	Partition tolerance (P)
	The system continues to operate despite an arbitrary number of messages being dropped
	(or delayed) by the network between nodes.

	Understanding CAP Trade-offs
	CA (Consistency + Availability, but no Partition Tolerance)
		Requires all nodes to communicate reliably.
		Example: Single-node relational databases (e.g., PostgreSQL, MySQL without replication).
	CP (Consistency + Partition Tolerance, but no Availability)
		Prioritizes consistency; some requests may fail during a partition.
		Example: MongoDB (when using strong consistency settings), HBase.
	AP (Availability + Partition Tolerance, but no Consistency)
		Prioritizes availability; data may be inconsistent temporarily.
		Example: DynamoDB, Cassandra, CouchDB.

	Why Does CAP Matter?
	In real-world distributed systems, network failures (partitions) are inevitable.
	So, databases must choose between Consistency (CP) or Availability (AP).
*/

// NOTE: Data hierarchy
/*
	Mongo -> Database -> Collection -> Document -> Fields
*/

// NOTE: Open mongo shell
/*
	mongosh -u <username> -p
	mongosh -u root -p
	 |-> prompt for password

	With connection string:
	mongosh "mongodb://<user name>:<password>@<dns or ip address>:<port>[/<database name>]"
	mongosh "mongodb://root:123456@localhost:27017"
*/

// NOTE: Close mongo shell
/*
	```mongo
	exit
	```
*/

// NOTE: Show current db
/*
	```mongo
	db
	```
*/

// NOTE: List all of the dbs
/*
	```mongo
	show dbs
	```

	If a database is empty and there is no collection in it, it won't be shown in the db list.
*/

// NOTE: Create a database
/*
	```mongo
	use <database_name>
	```
*/

// NOTE: Switch to a different database
/*
	```mongo
	use <database_name>
	```
*/

// NOTE: List collections in the database
/*
	1. show collections

	Command:
	```mongo
	show collections
	```

	Usage:
	This is a shell command in mongosh (MongoDB Shell).
	It lists all collections in the current database.
	Provides a simple, direct list of collection names.

	Example Output:
	```bash
	users
	orders
	products
	```

	2. db.getCollectionNames()

	Command:
	```mongo
	db.getCollectionNames()
	```

	Usage:
	This is a JavaScript method that returns an array of collection names.
	Useful when working inside scripts or programmatically retrieving collections.

	Example Output:
	```bash
	[ "users", "orders", "products" ]
	```

	Key Differences:
	Feature			show collections	db.getCollectionNames()
	Type			Shell command		JavaScript function
	Returns			Plain list			Array of strings
	Best for		Quick listing		Programmatic use
	Output format	Plain text			JavaScript array

*/

// NOTE: Check the stats of the current database
/*
	```mongo
	db.stats()
	```
*/

// NOTE: Insert document
/*
	Syntax:
	```mongo
	db.<collection_name>.insert(<object>)
	```

	Example:
	```mongo
	db.users.insert({"name": "Amir"})
	```
	This will create a users collection and insert this record into it.

	DeprecationWarning: Collection.insert() is deprecated. Use insertOne, insertMany, or bulkWrite.

	Syntax:
	```mongo
	db.<collection_name>.insertOne(<object>)
	```
*/

// NOTE: Insert one into document
/*
	Syntax:
	```mongo
	db.<collection_name>.insertOne(<object>)
	```
*/

// NOTE: Clear the screen
/*
	cls
	This is a built-in shortcut in mongosh that clears the screen.

	Syntax:
	```mongo
	cls
	```

	console.clear()
	This is a JavaScript method that works inside mongosh to clear the terminal output.

	Syntax:
	```mongo
	console.clear()
	```
*/

// NOTE: View documents
/*
	Syntax:
	```mongo
	db.<collection_name>.find()
	```

	Example:
	```mongo
	db.users.find()
	```

	Example:
	```mongo
	db.users.find().pretty()
	```
*/

// NOTE: Drop database
/*
	Syntax:
	```mongo
	db.dropDatabase()
	```
*/

func main() {

}
