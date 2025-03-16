package main

// NOTE: CHAR (Fixed-Length Character String)
/*
	Definition: CHAR(n) stores a fixed-length string of exactly n characters.
	If the input string is shorter than n, it is padded with spaces to reach the specified length.

	Storage: Always uses n bytes of storage, regardless of the actual length of the string.

	Use Case:
	 - When the length of the data is constant and predictable (e.g., country codes like 'US', 'CA', or 'IN').
	 - When you need to enforce a specific length for all entries.

	```sql
	CREATE TABLE example (
		country_code CHAR(2) NOT NULL
	);
	```

	'US' will be stored as 'US'.
	'A' will be stored as 'A ' (padded with a space).
*/

// NOTE: VARCHAR (Variable-Length Character String)
/*
	Definition: VARCHAR(n) stores a variable-length string of up to n characters.
	It only uses as much storage as needed for the actual string, plus a small overhead for storing the length.

	Storage: Uses storage proportional to the length of the string (plus 1-2 bytes for length metadata).

	Use Case:
	 - When the length of the data varies significantly (e.g., names, addresses, descriptions).
	 - When you want to save storage space by not padding shorter strings.

	Example:

	```sql
	CREATE TABLE example (
		name VARCHAR(50) NOT NULL
	);
	```

	'John' will be stored as 'John' (no padding).
	'Alice' will be stored as 'Alice'.
*/

// NOTE: Key Differences
/*
	Feature			CHAR(n)								VARCHAR(n)
	Length			Fixed length (n characters).		Variable length (up to n chars).
	Storage			Always uses n bytes.				Uses storage proportional to data.
	Padding			Pads with spaces if shorter.		No padding.
	Performance		Slightly faster for fixed lengths.	Slightly slower due to overhead.
	Use Case		Fixed-length data (e.g., codes).	Variable-length data (e.g., names).
*/

// NOTE: When to Use Which?
/*
	Use CHAR:
	 - When the data length is always the same (e.g., CHAR(2) for country codes, CHAR(10) for phone numbers in a specific format).
	 - When you want to enforce a fixed length for consistency.

	Use VARCHAR:
	 - When the data length varies (e.g., names, addresses, descriptions).
	 - When you want to save storage space by avoiding padding.
*/

// NOTE: Postgres field types
/*
	PostgreSQL supports a wide variety of data types (also called field types) to store different kinds of data.
	These data types are categorized into several groups based on their purpose. Below is a comprehensive explanation
	of the most commonly used data types in PostgreSQL:
*/

// NOTE: 1. Numeric Types
/*
	Used for storing numbers, both integers and floating-point values.

	Data Type			descriptions																		Example
	SMALLINT			2-byte integer, range: -32,768 to 32,767.											42
	INT or INTEGER		4-byte integer, range: -2,147,483,648 to 2,147,483,647.								123456
	BIGINT				8-byte integer, range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.		123456789012345
	DECIMAL(p, s)		Exact numeric, user-defined precision (p) and scale (s).							123.45
	NUMERIC(p, s)		Same as DECIMAL.																	123.45
	REAL				4-byte floating-point number, 6 decimal digits precision.							123.456
	DOUBLE PRECISION	8-byte floating-point number, 15 decimal digits precision.							123.456789012345
	SERIAL				Auto-incrementing 4-byte integer (used for primary keys).							1, 2, 3, ...
	BIGSERIAL			Auto-incrementing 8-byte integer (used for primary keys).							1, 2, 3, ...
*/

// NOTE: 2. Character Types
/*
	Used for storing text data.

	Data Type	Description										Example
	CHAR(n)		Fixed-length string, padded with spaces.		'ABC '
	VARCHAR(n)	Variable-length string, up to n characters.		'ABC'
	TEXT		Variable-length string, unlimited length.		'This is a long text.'
	NAME		Internal type for object names, 64 bytes max.	'my_table'
*/

// NOTE: 3. Binary Data Types
/*
	Used for storing binary data.

	Data Type	Description						Example
	BYTEA		Variable-length binary data.	'\xDEADBEEF'
*/

// NOTE: 4. Date/Time Types
/*
	Used for storing dates, times, and intervals.

	Data Type		Description										Example
	DATE			Stores a date (year, month, day).				'2023-10-05'
	TIME			Stores a time of day (hour, minute, second).	'14:30:00'
	TIMESTAMP		Stores both date and time.						'2023-10-05 14:30:00'
	TIMESTAMPTZ		Stores date, time, and time zone.				'2023-10-05 14:30:00+00'
	INTERVAL		Stores a time interval.							'1 day 2 hours 30 minutes'
*/

// NOTE: 5. Boolean Type
/*
	Used for storing true/false values.

	Data Type		Description						Example
	BOOLEAN			Stores TRUE, FALSE, or NULL.	TRUE
*/

// NOTE: 6. Geometric Types
/*
	Used for storing geometric shapes and points.

	Data Type	Description										Example
	POINT		A geometric point in 2D space.					(3.0, 4.0)
	LINE		An infinite line in 2D space.					{A, B, C}
	LSEG		A line segment in 2D space.						[(1,1), (2,2)]
	BOX			A rectangular box in 2D space.					(2,2), (3,3)
	PATH		A closed or open geometric path in 2D space.	[(1,1), (2,2), (3,3)]
	POLYGON		A closed geometric path in 2D space.			((1,1), (2,2), (3,3))
	CIRCLE		A circle in 2D space.							<(1,1), 5>
*/

// NOTE: 7. Network Address Types
/*
	Used for storing IP addresses and network information.

	Data Type		Description									Example
	INET			Stores an IPv4 or IPv6 address.				'192.168.1.1'
	CIDR			Stores an IPv4 or IPv6 network address.		'192.168.1.0/24'
	MACADDR			Stores a MAC address.						'08:00:2b:01:02:03'
*/

// NOTE: 8. JSON Types
/*
	Used for storing JSON data.

	Data Type		Description																Example
	JSON			Stores JSON data as plain text.											'{"key": "value"}'
	JSONB			Stores JSON data in a binary format (more efficient for querying).		'{"key": "value"}'
*/

// NOTE: 9. Array Types
/*
	Used for storing arrays of any data type.

	Data Type		descriptions			Example
	INT[]			Array of integers.		'{1, 2, 3}'
	TEXT[]			Array of text.			'{"a", "b", "c"}'
*/

// NOTE: 10. Composite Types
/*
	Used for defining custom row types.

	Data Type		Description						Example
	CREATE TYPE		User-defined composite type.	CREATE TYPE person AS (name TEXT, age INT);
*/

// NOTE: 11. Range Types
/*
	Used for storing a range of values.

	Data Type		Description				Example
	INT4RANGE		Range of integers.		[1, 10)
	TSRANGE			Range of timestamps.	['2023-10-01', '2023-10-31')
*/

// NOTE: 12. UUID Type
/*
	Used for storing universally unique identifiers (UUIDs).

	Data Type	Description			Example
	UUID		Stores a UUID.		'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'
*/

// NOTE: 13. Special Types
/*
	Used for specific purposes.

	Data Type		Description							Example
	MONEY			Stores currency amounts.			'$123.45'
	XML				Stores XML data.					'<note>Hello</note>'
	TSVECTOR		Used for full-text search.			'quick brown fox'
	TSQUERY			Used for full-text search queries.	'quick & brown'
*/

// NOTE: Summary
/*
	PostgreSQL provides a rich set of data types to handle various kinds of data,
	from simple numbers and text to complex geometric shapes and JSON documents.
	Choosing the right data type is crucial for optimizing storage and performance.
*/

func main() {

}
