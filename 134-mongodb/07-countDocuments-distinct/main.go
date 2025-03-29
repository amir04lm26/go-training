package main

// NOTE: countDocuments() Method
/*
	This method returns the count of documents matching a query. Unlike the deprecated .count(),
	it provides accurate results even with large datasets.

	Syntax:
	```mongo
	db.collection.countDocuments(query)
	```

	Example:
	```mongo
	db.users.countDocuments({ age: { $gte: 18 } }); // Count users who are 18 or older
	```
*/

// NOTE: distinct() Method
/*
	Retrieves unique values for a specified field across all documents in a collection.

	Syntax:
	```mongo
	db.collection.distinct("field", query)
	```

	Example:
	```mongo
	db.users.distinct("country"); // Get a list of unique countries from users
	db.users.distinct("city", { country: "USA" }); // Unique cities in the USA
	```
*/

func main() {

}
