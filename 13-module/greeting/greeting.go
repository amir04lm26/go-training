package greeting

import "fmt"

// NOTE: Each package should have a file with the same name as the package

func Hello(name string) string {
	message := fmt.Sprintf("Hi %s. Welcome!", name)
	return message
}
