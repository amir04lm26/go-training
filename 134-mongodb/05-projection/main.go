package main

// NOTE: Projection
/*
	Projection in MongoDB is used to control which fields are included or excluded in the query results.
	By default, MongoDB returns all fields of a document when querying a collection, but projection
	allows you to specify only the fields you need, improving performance by reducing data transfer.
*/

// NOTE: Syntax of Projection
/*
	Projection is specified as the second parameter in the find() method:

	```mongo
	db.collection.find({ query }, { projection })
	```

	1 → Includes the field (except _id, which is included by default unless explicitly excluded).
	0 → Excludes the field.
*/

// NOTE: Examples of Projection

// NOTE: 1. Including Specific Fields
/*
	Return only the name and age fields:

	```mongo
	db.users.find({}, { name: 1, age: 1, _id: 0 });
	```

	Result:
	```mongo
	[
		{ "name": "Alice", "age": 25 },
		{ "name": "Bob", "age": 30 }
	]
	```
*/

// NOTE: 2. Excluding Specific Fields
/*
	Exclude the age field:
	```mongo
	db.users.find({}, { age: 0 });
	```

	Result:
	```mongo
	[
	  { "_id": ObjectId("..."), "name": "Alice" },
	  { "_id": ObjectId("..."), "name": "Bob" }
	]
	```
*/

// NOTE: 3. Using Projection with Filters
/*
	Get only the name field for users whose age is greater than 25:
	```mongo
	db.users.find({ age: { $gt: 25 } }, { name: 1, _id: 0 });
	```

	Result:
	```mongo
	[
		{ "name": "Bob" }
	]
	```
*/

// NOTE: 4. Projection for Embedded Fields
/*
	For documents with embedded objects, you can project nested fields.

	Example Document:
	```mongo
	{
	  "_id": ObjectId("..."),
	  "name": "Alice",
	  "address": { "city": "New York", "zip": "10001" }
	}
	```

	Query:
	```mongo
	db.users.find({}, { "address.city": 1, _id: 0 });
	```

	Result:
	```mongo
	[
	  { "address": { "city": "New York" } }
	]
	```
*/

// NOTE: When to Use Projection?
/*
	To reduce the amount of data transferred from the database to the application.
	To improve query performance by retrieving only necessary fields.
	To limit sensitive data exposure (e.g., excluding passwords or personal details).
*/

func main() {

}
