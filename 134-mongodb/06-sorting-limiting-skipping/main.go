package main

// NOTE: Sorting & Limiting Methods in MongoDB
/*
	These methods control the number of documents retrieved and their order.
	Note that he order of .sort() and .limit() matters in MongoDB.
*/

// NOTE: 1. .sort()
/*
	Used to arrange query results in ascending or descending order based on one or more fields.

	Syntax:
	```mongo
	db.collection.find().sort({ field: 1 })  // Ascending
	db.collection.find().sort({ field: -1 }) // Descending
	db.collection.find().sort({ age: 1, name: -1 }) // Sort by multiple fields
	```
	Example:
	```mongo
	db.users.find().sort({ age: -1 }); // Sort users by age in descending order
	```
*/

// NOTE: 2. .limit()
/*
	Restricts the number of documents returned.

	Syntax:
	```mongo
	db.collection.find().limit(N)
	```

	Example:
	```mongo
	db.users.find().limit(5); // Returns only 5 users
	```
*/

// NOTE: 3. .skip()
/*
	Skips a certain number of documents before returning the result.

	Syntax:
	```mongo
	db.collection.find().skip(N)
	```

	Example:
	```mongo
	db.users.find().skip(10).limit(5); // Skips 10 documents, then returns 5 users
	```
*/

func main() {

}
