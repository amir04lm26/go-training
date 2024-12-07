package main

import "fmt"

// NOTE: Dependency Injection in Go
/*
	Go does not have a native dependency injection framework, but you can achieve it by passing
	dependencies through function parameters or struct fields.
*/

// NOTE: Example: Manual Dependency Injection

// NOTE: Explanation:
/*
	•	Interface-based dependency: The Service struct depends on the Database interface,
		allowing different implementations (like MySQL) to be injected.
	•	Constructor-like pattern: The dependencies are passed through struct fields or
		function parameters.
*/

// * Database interface
type Database interface {
	Query() string
}

// * MySQL implementation of Database
type MySQL struct{}

func (db *MySQL) Query() string {
	return "MySQL query result"
}

// * Service that uses Database
type Service struct {
	Database Database
}

func (service *Service) GetData() string {
	return service.Database.Query()
}

func main() {
	db := &MySQL{}
	service := &Service{Database: db}
	fmt.Println(service.GetData())
}
