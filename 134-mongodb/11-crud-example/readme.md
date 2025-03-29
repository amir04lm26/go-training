1. bson.D (Ordered Document)
   bson.D is a slice of key-value pairs and preserves field order.
   It's useful when field order matters (e.g., for $set operations or pipeline stages in aggregation).

Example:

```mongo
filter := bson.D{
    {"author", "H. G. Wells"},
    {"price", bson.D{{"$gte", 10}}},
}
```

✅ Preserves order
✅ Recommended for queries where field order is important
❌ Slightly more verbose than bson.M

2. bson.M (Unordered Map)
   bson.M is a map of key-value pairs and does not preserve field order.
   It's more convenient for simple key-value lookups but should not be used when order matters.

Example:

```mongo
filter := bson.M{
    "author": "H. G. Wells",
    "price": bson.M{"$gte": 10},
}
```

✅ Easier to write (similar to JSON or Go maps)
✅ Ideal for simple queries where order doesn’t matter
❌ Does not guarantee order (could cause issues in some queries like $set)

Key Differences Table
Feature bson.D (Ordered) bson.M (Unordered)
Type Slice of structs ([]bson.E) Map (map[string]interface{})
Order Preserves order Does not preserve order
Use Case When order matters (e.g., $set, aggregation stages) When order doesn’t matter (e.g., filtering)
Readability More structured, but verbose Shorter and simpler
Performance Slightly slower (due to slice operations) Faster for direct lookups
