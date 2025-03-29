package main

// NOTE: 1. Create an Admin User (Superuser)
/*
	Admin users are created to manage the MongoDB database.
	First, connect to the MongoDB shell and create an admin user.

	Create Admin User:
	```mongo
	use admin;  // Switch to the 'admin' database

	db.createUser({
		user: "admin",
		pwd: "adminPassword123",  // Choose a strong password
		roles: [
			{ role: "root", db: "admin" }  // 'root' role allows full access to all databases
		]
	});
	```
*/

// NOTE: 2. Create a User with Read-Write Access
/*
	For regular application users, you can assign them specific roles such as readWrite on a database.

	Create a Database User with Read-Write Access:
	```mongo
	use myDatabase;  // Switch to the database where the user will be created

	db.createUser({
		user: "appUser",
		pwd: "appUserPassword123",  // Choose a strong password
		roles: [
			{ role: "readWrite", db: "myDatabase" }  // Allows read and write access on 'myDatabase'
		]
	});
	```
*/

//NOTE: 3. Create a User with Read-Only Access
/*
	If a user only needs to read from a database (but not modify anything), you can assign the
	read role.

	Create a Read-Only User:
	```mongo
	use myDatabase;

	db.createUser({
		user: "readOnlyUser",
		pwd: "readOnlyPassword123",  // Choose a strong password
		roles: [
			{ role: "read", db: "myDatabase" }  // Allows read access only on 'myDatabase'
		]
	});
	```
*/

//NOTE: 4. Create a User with Specific Roles (Custom Roles)
/*
	MongoDB allows you to assign specific custom roles to users, such as dbAdmin or clusterAdmin.

	Create a User with Database Admin Role:
	```mongo
	use myDatabase;

	db.createUser({
		user: "dbAdminUser",
		pwd: "dbAdminPassword123",  // Choose a strong password
		roles: [
			{ role: "dbAdmin", db: "myDatabase" }  // Grants admin privileges for managing the database
		]
	});
	```
*/

// NOTE: 5. Create a User with Multiple Roles
/*
	A user can have multiple roles depending on the access required.

	Create a User with Multiple Roles:
	```mongo
	use myDatabase;

	db.createUser({
		user: "appUserWithMultipleRoles",
		pwd: "appUserWithMultipleRolesPassword123",  // Choose a strong password
		roles: [
			{ role: "readWrite", db: "myDatabase" },  // Read and write access on 'myDatabase'
			{ role: "dbAdmin", db: "myDatabase" }  // Database admin role on 'myDatabase'
		]
	});
	```
*/

// NOTE: 6. Create a User with Cluster Admin Role (For Managing Sharded Clusters)
/*
	If you are working with a sharded cluster, a user with clusterAdmin role is required for
	cluster-level administration.

	Create a Cluster Admin User:
	```mongo
	use admin;

	db.createUser({
		user: "clusterAdminUser",
		pwd: "clusterAdminPassword123",  // Choose a strong password
		roles: [
			{ role: "clusterAdmin", db: "admin" }  // Grants access to manage the cluster
		]
	});
	```
*/

// NOTE: 7. Check Created Users
/*
	To list all users in a database:
	```mongo
	use myDatabase;  // Switch to the database

	show users
	or
	db.getUsers();   // List all users for that database
	```

	To list users across all databases:
	```mongo
	use admin;
	db.system.users.find();  // List all users in the MongoDB cluster
	```
*/

// NOTE: 8. Authentication and Authorization
/*
	Once the users are created, you must enable authentication for MongoDB to enforce user-based
	access control. To enable authentication:

	Enable Authentication in mongod.conf: Set security.authorization to "enabled" in the MongoDB
	configuration file (mongod.conf).
	```yaml
	security:
		authorization: "enabled"
	```

	Restart MongoDB: After modifying the configuration, restart MongoDB:
	```bash
		sudo systemctl restart mongod
		or
		brew services restart mongodb-community
	```

	Authenticate Using MongoDB User: Use the created user for authentication:
	```mongo
		db.auth("admin", "adminPassword123");
	```
*/

// NOTE: 9. Delete a User
/*
	To remove a user from a specific database:
	```mongo
		use myDatabase;

		db.dropUser("appUser");  // Drop the user 'appUser'
	```

	To remove a user from the entire MongoDB system:
	```mongo
	use admin;

	db.dropUser("admin");  // Drop the 'admin' user
	```
*/

// NOTE: Conclusion
/*
	MongoDB supports various user creation options with different levels of permissions (e.g.,
	readWrite, dbAdmin, root).
	You can assign users specific roles based on their responsibilities.
	Authentication must be enabled in the MongoDB configuration to enforce user roles.
*/

// NOTE: Use db.runCommand() to Get Current User
/*
	You can use the db.runCommand() with the connectionStatus command to get information about
	the current user.
	```mongo
	db.runCommand({ connectionStatus: 1 });
	```

	This command will return a document with details about the current user and their roles.
	Here’s an example of the output:
	```mongo
	{
		"authInfo": {
			"authenticatedUsers": [
				{
					"user": "appUser",
					"db": "myDatabase"
				}
			],
			"authenticatedUserRoles": [
				{ "role": "readWrite", "db": "myDatabase" }
			]
		},
		"ok": 1
	}
	```

	In this example:
	"user" represents the username of the currently authenticated user.
	"db" represents the database where the user was authenticated.
	"authenticatedUserRoles" shows the roles assigned to that user.
*/

// NOTE: db.getUser()
/*
	If you know the database, you can also use db.getUser() to get details about the user in
	that specific database.

	```mongo
	use myDatabase;  // Switch to the database
	db.getUser("appUser");
	```

	This will return the details of the user appUser in the myDatabase.
*/

// NOTE: Summary Table of MongoDB Roles
/*
	Role Name				Description
	read					Provides read-only access to all collections in a database.
	readWrite				Provides both read and write access to all collections in a database.
	dbAdmin					Administrative privileges for managing indexes and database statistics.
	dbOwner					Full access including the ability to manage users and roles in a database.
	userAdmin				Ability to create and modify users for the database.
	userAdminAnyDatabase	Ability to create and modify users across all databases.
	readAnyDatabase			Read-only access to all databases.
	readWriteAnyDatabase	Read and write access to all databases.
	dbAdminAnyDatabase		Administrative privileges for all databases.
	root					Full access to all databases and administrative commands.
	clusterAdmin			Cluster-level admin privileges (e.g., manage sharding, replication).
	clusterManager			Manage cluster configuration but not full admin rights.
	clusterMonitor			Read access to cluster statistics.
	hostManager				Manage hosts in a replica set or sharded cluster.
	backup					Privileges to back up the database.
	restore					Privileges to restore data from a backup.
	audit					Access to audit logs.
	manageDbUsers			Ability to create, modify, and delete users for a database.
	manageRoles				Ability to create, modify, and delete roles within the database.
	applicationUser			Basic application-level user access.

	These roles allow fine-grained access control to MongoDB resources, ensuring that users have
	the necessary privileges for their tasks while restricting unnecessary access. You can
	combine these roles or create custom roles to meet specific needs.
*/

func main() {

}
