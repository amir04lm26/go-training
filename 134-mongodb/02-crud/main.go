package main

// NOTE: Mongodb CRUD methods

// NOTE: 1. INSERT METHODS

// NOTE: insertOne()
/*
	The insertOne() method inserts a single document into a MongoDB collection.
	It is atomic and ensures that the document is added successfully.

	Usage Example:
	```sql
	db.users.insertOne({ name: "Alice", age: 25 });
	```

	- This adds a single document { name: "Alice", age: 25 } to the users collection.
	- Returns an acknowledgment with the inserted ID.
*/

// NOTE: insertMany()
/*
	The insertMany() method inserts multiple documents in a single operation.
	It allows bulk insertion and can be configured to continue inserting even
	if one document fails.

	Usage Example:
	```sql
	db.users.insertMany([
	{ name: "Bob", age: 30 },
	{ name: "Charlie", age: 35 }
	]);
	```

	- Inserts multiple documents efficiently.
	- By default, it is ordered, meaning if an error occurs, the remaining documents won’t be inserted.
	- You can set { ordered: false } to continue inserting the remaining documents even if an error
	  occurs.
*/

// NOTE: bulkWrite()
/*
	- Purpose: Performs multiple operations (insert, update, delete, replace) in a single batch.
	- Operations Supported: Insert, update, delete, replace.

	- Example:
	```sql
	db.users.bulkWrite([
		{ insertOne: { document: { name: "Charlie", age: 35 } } },
		{ updateOne: { filter: { name: "Alice" }, update: { $set: { age: 26 } } } },
		{ deleteOne: { filter: { name: "Bob" } } }
	]);
	```

	- Options:
		- ordered: true (default) → Stops on first error.
		- ordered: false → Continues even if some fail.
	- Use Case: When performing multiple types of operations in bulk.
*/

// NOTE: bulkWrite() (for Inserts)
/*
	The bulkWrite() method allows batch execution of inserts, updates, and deletes in a
	single operation. It is useful for mixed operations where you need to insert, update,
	and delete in a single request.

	Usage Example:
	```sql
	db.users.bulkWrite([
		{ insertOne: { document: { name: "David", age: 40 } } }
	]);
	```

	- Supports mixed operations (not just inserts).
	- It is more flexible than insertMany(), but requires additional syntax.
*/

// NOTE: Insert Methods Table
/*
	Method			Inserts Multiple?	Ordered Execution	Stops on Error?		  Use Case
	insertOne()		❌ No				N/A					N/A					 Insert a single document efficiently.
	insertMany()	✅ Yes				✅ Yes (default)		✅ Yes (default)		Bulk insert of multiple documents.
	bulkWrite()		✅ Yes				✅ Yes (default)		✅ Yes (default)		Perform mixed bulk operations efficiently.
*/

// NOTE: 2. UPDATE METHODS

// NOTE: updateOne()
/*
	The updateOne() method updates the first document that matches a given filter.
	If no document matches, it does nothing unless { upsert: true } is used.

	Usage Example:
	```sql
	db.users.updateOne({ name: "Alice" }, { $set: { age: 26 } });
	```

	- Updates the first document where name is "Alice".
	- If no document is found, nothing happens unless { upsert: true } is used.
*/

// NOTE: updateMany()
/*
	The updateMany() method updates all documents that match the given filter.
	It is useful for bulk updates.

	Usage Example:
	```sql
	db.users.updateMany({ age: { $gt: 30 } }, { $set: { status: "active" } });
	```

	- Updates all users where age > 30 and sets their status to "active".
*/

// NOTE: bulkWrite() (for Updates)
/*
	The bulkWrite() method can also perform updates, including both updateOne and updateMany.

	Usage Example:
	```sql
	db.users.bulkWrite([
		{ updateOne: { filter: { name: "Bob" }, update: { $set: { age: 31 } } } }
	]);
	```

	- Can execute multiple update operations in a single request.
*/

// NOTE: Update Methods Table
/*
	Method			Updates Multiple?	Returns Matched Count?		Use Case
	updateOne()		❌ No				✅ Yes						Update a single document efficiently.
	updateMany()	✅ Yes				✅ Yes						Modify multiple documents at once.
	bulkWrite()		✅ Yes				✅ Yes						Perform large-scale modifications efficiently.
*/

// NOTE: 3. DELETE METHODS

// NOTE: deleteOne()
/*
	The deleteOne() method removes the first document that matches the filter criteria.

	Usage Example:
	```sql
	db.users.deleteOne({ name: "Alice" });
	```

	- Deletes the first document where name = "Alice".
*/

// NOTE: deleteMany()
/*
	The deleteMany() method deletes all documents that match a filter.

	Usage Example:
	```sql
	db.users.deleteMany({ age: { $lt: 20 } });
	```

	- Deletes all users whose age is less than 20.
*/

// NOTE: bulkWrite() (for Deletes)
/*
	The bulkWrite() method can also perform delete operations in bulk.

	Usage Example:
	```sql
	db.users.bulkWrite([
		{ deleteOne: { filter: { name: "Bob" } } }
	]);
	```

	- Useful when deleting multiple documents efficiently.
*/

// NOTE: Delete Methods Table
/*
	Method			Deletes Multiple?	Returns Deleted Count?	 Use Case
	deleteOne()		❌ No				✅ Yes					Remove a single document.
	deleteMany()	✅ Yes				✅ Yes					Remove multiple documents at once.
	bulkWrite()		✅ Yes				✅ Yes					Perform multiple deletions efficiently.
*/

// NOTE: 4. QUERY METHODS

// NOTE: findOne()
/*
	The findOne() method retrieves only the first document that matches a query.

	Usage Example:
	```sql
	db.users.findOne({ name: "Alice" });
	```

	- Returns the first document where name = "Alice".
*/

// NOTE: find()
/*
	The find() method retrieves all documents that match a filter.

	Usage Example:
	```sql
	db.users.find({ age: { $gt: 25 } });
	```

	- Returns all users where age > 25.
*/

// NOTE: aggregate()
/*
	The aggregate() method allows performing advanced operations like grouping, sorting, and filtering.

	Usage Example:
	```sql
	db.users.aggregate([
		{ $match: { age: { $gt: 25 } } },
		{ $group: { _id: "$status", total: { $sum: 1 } } }
	]);
	```

	- Groups users by status and counts the total in each group.
*/

// NOTE: Query Methods Table
/*
	Method			Returns All Matches?		Limits Output?					 Use Case
	findOne()		❌ No (first match only)		✅ Yes							Retrieve a single document quickly.
	find()			✅ Yes						❌ No (unless .limit() is used)	Retrieve multiple documents.
	aggregate()		✅ Yes						❌ No (unless $limit is used)	Data transformation and analytics.
*/

// NOTE: Final Thoughts
/*
	Use insertOne() or insertMany() for adding documents.
	Use updateOne() or updateMany() for modifying documents.
	Use deleteOne() or deleteMany() for removing documents.
	Use findOne(), find(), or aggregate() for reading data.
	Use bulkWrite() when performing multiple operations efficiently.
*/

// NOTE: Other Methods

// NOTE: replaceOne():
/*
	Purpose: Completely replaces a single document with a new one.

	How it works: Finds a document that matches the filter and replaces it entirely with
	the new document. The _id field is not automatically preserved unless explicitly
	included in the new document.

	Example:
	```mongo
		db.collection.replaceOne(
			{ _id: 1 },  // filter condition
			{ name: "Alice", age: 26 }  // new document to replace the existing one
		);
	```
*/

// NOTE: findOneAndUpdate()
/*
	Purpose: Finds a document, applies an update, and returns the document before or after the update.

	How it works: It combines a find and update operation, returning either the document before or
	after the update. This method allows fine-grained control over whether the old or new document
	is returned.

	Example:
	```mongo
		db.collection.findOneAndUpdate(
			{ _id: 1 },
			{ $set: { age: 26 } },
			{ returnDocument: "after" }
		);
	```
*/

// NOTE: findOneAndReplace()
/*
	Purpose: Finds a document, replaces it completely, and returns the document before or
	after replacement.

	How it works: Similar to replaceOne(), but this method allows you to get the document
	before or after it is replaced.

	Example:
	```mongo
	db.collection.findOneAndReplace(
		{ _id: 1 },
		{ name: "Alice", age: 26 },
		{ returnDocument: "after" }
	);
	```
*/

// NOTE: findOneAndDelete()
/*
	Purpose: Finds and deletes a single document.

	How it works: Deletes the first document that matches the filter and returns the deleted document.

	Example:
	```mongo
	db.collection.findOneAndDelete(
		{ _id: 1 },
		{ projection: { _id: 0 } }
	);
	```
*/

// NOTE: upsert (used with updateOne() and updateMany())
/*
	Purpose: Performs an update if the document exists, or inserts a new document if it does not.

	How it works: By setting the upsert option to true, MongoDB will either insert a new document
	if none matches the filter, or update an existing document.

	Example:
	```sql
	db.collection.updateOne(
		{ _id: 1 },
		{ $set: { age: 26 } },
		{ upsert: true }
	);
	```
*/

func main() {

}
