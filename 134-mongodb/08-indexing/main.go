package main

// NOTE: Indexing Methods in MongoDB
/*
	Indexes improve query performance by optimizing searches, sorting, and filtering.
*/

// NOTE: 1. .createIndex()
/*
	Creates an index on a field to speed up searches.

	Syntax:
	```mongo
	db.collection.createIndex({ field: 1 })  // Ascending index
	db.collection.createIndex({ field: -1 }) // Descending index
	```

	Example:
	```mongo
	db.users.createIndex({ age: 1 }); // Creates an index on the `age` field
	db.users.createIndex({ name: 1, age: -1 }); // Multi-field index
	```
*/

// NOTE: 2. .getIndexes()
/*
	Lists all indexes on a collection.

	Syntax:
	```mongo
	db.collection.getIndexes()
	```

	Example:
	```mongo
	db.users.getIndexes(); // Show all indexes in the users collection
	```
*/

// NOTE: 3. .dropIndex()
/*
	Removes an existing index.

	Syntax:
	```mongo
	db.collection.dropIndex("index_name")
	```

	Example:
	```mongo
	db.users.dropIndex({ age: 1 }); // Drops index on `age`
	```
*/

func main() {

}
