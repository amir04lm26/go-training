package main

// NOTE: Operators
/*
	MongoDB provides a wide range of operators for performing various operations like querying,
	updating, and aggregating data.
*/

// NOTE: 1. Query Operators
/*
	Query operators are used to specify criteria when querying documents in MongoDB.
*/

// NOTE: Comparison Operators:
/*
	$eq: Matches values that are equal to a specified value.
	```mongo
	db.collection.find({ age: { $eq: 30 } })
	```

	$gt: Matches values that are greater than the specified value.
	```mongo
	db.collection.find({ age: { $gt: 30 } })
	```

	$gte: Matches values that are greater than or equal to the specified value.
	```mongo
	db.collection.find({ age: { $gte: 30 } })
	```

	$lt: Matches values that are less than the specified value.
	```mongo
	db.collection.find({ age: { $lt: 30 } })
	```

	$lte: Matches values that are less than or equal to the specified value.
	```mongo
	db.collection.find({ age: { $lte: 30 } })
	```

	$ne: Matches values that are not equal to the specified value.
	```mongo
	db.collection.find({ age: { $ne: 30 } })
	```

	$in: Matches any value within the specified array.
	```mongo
	db.collection.find({ age: { $in: [20, 25, 30] } })
	```

	$nin: Matches values that are not in the specified array.
	```mongo
	db.collection.find({ age: { $nin: [20, 25, 30] } })
	```
*/

// NOTE: Logical Operators:
/*
	$and: Joins multiple query conditions with logical AND.
	```mongo
	db.collection.find({ $and: [{ age: { $gte: 30 } }, { name: "John" }] })
	```

	$or: Joins multiple query conditions with logical OR.
	```mongo
	db.collection.find({ $or: [{ age: { $gte: 30 } }, { name: "John" }] })
	```

	$not: Inverts the effect of a query expression.
	```mongo
	db.collection.find({ age: { $not: { $gte: 30 } } })
	```

	$nor: Joins multiple query conditions with logical NOR (not OR).
	```mongo
	db.collection.find({ $nor: [{ age: { $gte: 30 } }, { name: "John" }] })
	```
*/

// NOTE: Element Operators:
/*
	$exists: Matches documents that have (or do not have) the specified field.
	```mongo
	db.collection.find({ age: { $exists: true } })
	```

	$type: Matches documents where the value of a field is of the specified BSON data type.
	```mongo
	db.collection.find({ age: { $type: "int" } })
	```
*/

// NOTE: Array Operators:
/*
	$all: Matches documents where the field is an array that contains all the specified elements.
	```mongo
	db.collection.find({ tags: { $all: ["tag1", "tag2"] } })
	```

	$elemMatch: Matches documents that contain an array field with at least one element that matches
	the specified query.
	```mongo
	db.collection.find({ tags: { $elemMatch: { type: "A", value: 10 } } })
	```

	$size: Matches documents where the field is an array with a specified number of elements.
	```mongo
	db.collection.find({ tags: { $size: 3 } })
	```
*/

// NOTE: 2. Update Operators
/*
	Update operators modify the documents in a collection.
*/

// NOTE: Field Update Operators:
/*
	$set: Sets the value of a field. If the field does not exist, it is created.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $set: { age: 30 } })
	```

	$unset: Removes a field from a document.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $unset: { age: "" } })
	```

	$inc: Increments a field by a specified value.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $inc: { age: 1 } })
	```

	$mul: Multiplies the value of a field by a specified value.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $mul: { price: 2 } })
	```

	$rename: Renames a field.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $rename: { oldField: "newField" } })
	```
*/

// NOTE: Array Update Operators:
/*
	$push: Adds an element to an array.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $push: { tags: "newTag" } })
	```

	$push with $each: Adds multiple elements to an array.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $push: { tags: { $each: ["tag1", "tag2"] } } })
	```

	$addToSet: Adds an element to an array only if it does not already exist in the array (like a set).
	```mongo
	db.collection.updateOne({ _id: 1 }, { $addToSet: { tags: "uniqueTag" } })
	```

	$pop: Removes the first or last element of an array.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $pop: { tags: 1 } })  // removes last element

	$pop: { field: 1 }: Removes the last element from the array.

	$pop: { field: -1 }: Removes the first element from the array.
	```

	$pull: Removes all instances of a value from an array.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $pull: { tags: "unwantedTag" } })
	```

	$pullAll: Removes all specified values from an array.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $pullAll: { tags: ["tag1", "tag2"] } })
	```

	$slice: Limits the number of elements in an array to a specified size.
	```mongo
	db.collection.updateOne({ _id: 1 }, { $slice: { tags: 3 } })
	```
*/

// NOTE: 3. Aggregation Operators
/*
	Aggregation operators are used to perform operations such as filtering, sorting, and grouping data.

	$match: Filters the documents that enter the aggregation pipeline.
	```mongo
	db.collection.aggregate([{ $match: { age: { $gte: 30 } } }])
	```

	$group: Groups documents by a specified identifier and can include accumulators like $sum, $avg, etc.
	```mongo
	db.collection.aggregate([{ $group: { _id: "$age", total: { $sum: 1 } } }])
	```

	$sort: Sorts the documents by a specified field.
	```mongo
	db.collection.aggregate([{ $sort: { age: 1 } }])  // ascending order
	```

	$project: Shapes the documents to include only the specified fields.
	```mongo
	db.collection.aggregate([{ $project: { name: 1, age: 1 } }])
	```

	$limit: Limits the number of documents that enter the pipeline.
	```mongo
	db.collection.aggregate([{ $limit: 5 }])
	```

	$skip: Skips the first N documents in the pipeline.
	```mongo
	db.collection.aggregate([{ $skip: 5 }])
	```

	$unwind: Deconstructs an array field from the input documents to output a document
	for each element in the array.
	```mongo
	db.collection.aggregate([{ $unwind: "$tags" }])
	```

	Example:
	```mongo
	db.users.find()

	out:
	[
		{ "_id": 1, "name": "Alice", "hobbies": ["reading", "cycling", "hiking"] },
		{ "_id": 2, "name": "Bob", "hobbies": ["photography", "painting"] },
		{ "_id": 3, "name": "Charlie", "hobbies": ["swimming", "running"] }
	]

	db.users.aggregate([
		{ $unwind: "$hobbies" }
	])

	out:
	[
		{ "_id": 1, "name": "Alice", "hobbies": "reading" },
		{ "_id": 1, "name": "Alice", "hobbies": "cycling" },
		{ "_id": 1, "name": "Alice", "hobbies": "hiking" },
		{ "_id": 2, "name": "Bob", "hobbies": "photography" },
		{ "_id": 2, "name": "Bob", "hobbies": "painting" },
		{ "_id": 3, "name": "Charlie", "hobbies": "swimming" },
		{ "_id": 3, "name": "Charlie", "hobbies": "running" }
	]
	```

	Complex Example:
	```mongo
	db.users.aggregate([{
		$match: {name: {$exists: true}}
	},{
		$group: {_id: "$name", total: {$sum: 1}}
	}, {
		$sort: {total: -1}
	}])
	```

	Also we can use variables inside sum or avg
	```
	$group: {_id: "$name", total: {$sum: "$age"}}
	```
*/

// NOTE: 4. Text Search Operators
/*
	MongoDB also supports text search operators, which allow for querying text indexes.

	$text: Performs a text search on a collection using a text index.
	```mongo
	db.collection.find({ $text: { $search: "keyword" } })
	```

	$language: Specifies the language for text search.
	```mongo
	db.collection.find({ $text: { $search: "keyword", $language: "english" } })
	```

	$caseSensitive: Determines whether the text search is case-sensitive.
	```mongo
	db.collection.find({ $text: { $search: "Keyword", $caseSensitive: true } })
	```
*/

func main() {
}
