package main

// NOTE: Go application’s directory structure
/*
	Organizing a Go application’s directory structure effectively can make your code more readable,
	maintainable, and scalable. While Go doesn’t enforce a specific structure, developers follow
	some common patterns based on the nature of the project (monoliths, microservices, CLI tools,
	libraries, etc.). Below are some recommended structures depending on your project type.
*/

// NOTE:  1. Standard Directory Structure for Go Applications
/*
	This structure is inspired by community best practices (e.g., Go Project Layout):
	/your-app
	│
	├── cmd/                 # Command-line tools / application entry points
	│   ├── your-app/        # Main app entry (main.go)
	│   └── your-cli/        # (Optional) CLI tool entry (main.go)
	│
	├── pkg/                 # Reusable packages (importable by other projects)
	│   └── utils/           # Example utility package (e.g., helpers, validators)
	│
	├── internal/            # Private packages used only within the project
	│   ├── db/              # Database-related logic (e.g., migrations, queries)
	│   └── auth/            # Authentication and authorization logic
	│
	├── api/                 # API-related logic (handlers, routes, controllers)
	│   ├── handlers.go
	│   └── middlewares.go
	│
	├── configs/             # Configuration files (YAML, JSON, .env)
	│   └── config.yaml
	│
	├── migrations/          # Database migration files (SQL scripts)
	│
	├── web/                 # Static web assets (if applicable)
	│   └── templates/       # HTML templates
	│
	├── tests/               # Integration or E2E tests
	│   └── main_test.go
	│
	├── vendor/              # Vendor dependencies (if using `go mod vendor`)
	│
	├── go.mod               # Go module definition file
	├── go.sum               # Go dependencies checksum
	└── README.md            # Project documentation
*/

// NOTE: Explanation of Key Directories

// NOTE: cmd/
/*
	•	This folder contains entry points for your application, typically main.go files.
	•	For example, if your application has multiple binaries, you can create subfolders within cmd/.
*/

// NOTE: Example
/*
	#cmd/your-app/main.go
	package main

	import "your-app/internal/app"

	func main() {
		app.Run()
	}
*/

// NOTE: 2.	pkg/
/*
	•	This is where reusable and shareable packages reside. Other projects can import these
		packages if required.
	•	Example: pkg/utils might contain helper functions like logging or input validation.
*/

// NOTE: 3.	internal/
/*
	•	Contains project-specific packages that should not be imported by other projects.
	•	This is enforced by Go (packages in internal/ are not importable from outside the project).
*/

// NOTE: 4.	api/
/*
	•	API-related logic, such as HTTP handlers, middleware, and routing definitions.
*/

// NOTE: Example
/*
	# api/handlers.go
	package api

	import (
		"net/http"
	)

	func HelloHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
*/

// NOTE: 5.	configs/
/*
	•	Stores configuration files like YAML, JSON, or environment variables that are needed
		to run the app.
*/

// NOTE: 6.	tests/
/*
	•	Contains integration and end-to-end tests. Unit tests typically live next to the package
		they test (e.g., auth_test.go in internal/auth).
*/

// NOTE: 7.	migrations/
/*
	•	Database migration files if your application requires schema migrations.
*/

// NOTE: 8.	go.mod / go.sum
/*
	•	These files manage dependencies for your project using Go modules.
*/

// NOTE: 2. Simplified Structure for Smaller Applications
/*
	If your application is small and doesn’t require a complex structure, you can simplify it:

	/your-app
	│
	├── main.go            # Entry point for the application
	├── api/               # API-related code (handlers, routes)
	├── config/            # Configuration files
	├── db/                # Database logic
	├── utils/             # Helper functions or utilities
	├── go.mod             # Go module definition
	└── go.sum             # Go dependencies checksum
*/

// NOTE: 3. Example Structure for Microservices
/*
	In a microservice-oriented system, each service can have a self-contained structure:

	/user-service
	├── cmd/user-service/   # Service entry point (main.go)
	├── internal/           # Internal packages (auth, models, etc.)
	├── api/                # API handlers and routes
	├── db/                 # Database migrations and queries
	├── configs/            # Configuration files (YAML, env)
	├── go.mod              # Go module definition
	└── go.sum              # Go dependencies checksum
*/

// NOTE: 4. Best Practices for Go Project Structure
/*
	•	Use Go modules: Always initialize your project with go mod init <module-name>.
	•	Keep packages small and cohesive: Each package should have a single responsibility.
	•	Prefer internal/ for private packages: This helps to avoid leaking internal details.
	•	Write tests close to the code: Place unit tests alongside the package they test.
	•	Avoid deeply nested directories: Keep your structure simple and easy to navigate.
*/
