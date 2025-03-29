package main

// NOTE: Create implicitly
/*
	```sql
	db.<collection_name>.insert({"name": "Mcleod"})
	```
*/

// NOTE: Create explicitly
/*
	```sql
	db.createCollection(<name>, { <optional options> })
	```

	Example:
	```sql
	db.createCollection("customers")

	db.createCollection("crs", {capped: true, size: 65536, max: 1000000})
	```

	Advanced Example:
	```sql
	db.createCollection("users", {
	capped: true,
	size: 1024,
	max: 100,
	validator: {
		$jsonSchema: {
			bsonType: "object",
			required: ["name", "email"],
			properties: {
				name: { bsonType: "string", description: "must be a string and is required" },
				email: { bsonType: "string", pattern: "^.+@.+$", description: "must be a valid email" }
			}
		}
	}
	})
	```

	Another Example:
	```sql
	db.createCollection("products", {
		collation: { locale: "fr", strength: 2 }
	})
	```

	- This ensures French language sorting rules apply to the collection.
*/

// NOTE: `createCollection` Optional Parameters
/*
	Option						Type		Description
	capped						Boolean		Set to true to create a capped collection (fixed-size).
	size						Number		Maximum size (in bytes) of a capped collection.
	max							Number		Maximum number of documents in a capped collection.
	autoIndexId (deprecated)	Boolean		Automatically creates _id index (deprecated in MongoDB 4.2+).
	storageEngine				Object		Specifies storage engine options.
	validator					Object		Defines validation rules for documents.
	validationLevel				String		"off", "moderate", or "strict" (how strictly validation is enforced).
	validationAction			String		"warn" (logs warnings) or "error" (rejects invalid documents).
	collation					Object		Specifies language-specific sorting rules.
	expireAfterSeconds			Number		Used with TTL indexes to expire documents after a certain time.
*/

// NOTE: NOTE
/*
	In most cases, MongoDB creates collections automatically when you insert a document.
	Use db.createCollection() only when you need specific configurations (like capped collections,
	validation, etc.).
*/

// NOTE: View collections
/*
	```sql
	show collections
	```
*/

// NOTE: Drop a collection
/*
	```sql
	db.<collection_name>.drop()
	```
*/

func main() {

}
