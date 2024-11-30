package main

func main() {
	var variable interface{} = "hello"

	// NOTE: Type assertion to check if i is a string
	stringValue, ok := variable.(string)
	if ok {
		println("'variable' is string:", stringValue)
	} else {
		println("'variable' is not a string.")
	}
}
