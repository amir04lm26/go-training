package greeting

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %s. Welcome!", name)
	return message
}
