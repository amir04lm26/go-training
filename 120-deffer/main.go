package main

// NOTE: deffer
/*
	deferred functions won't run when os.Exit runs
	all deferred functions will run after panic, and panicking propagates to the top of call-stack
	in reverse order
*/

// NOTE: Deffer 3 rules
/*
	1. A deferred function arguments are evaluated when the defer statement is evaluated
	```go
		func a() {
			i := 0
			defer fmt.Println(i) // * output: 0
			i++
			return
		}
	```

	2.Deferred function calls are executed in Last In First Out order after the surrounding function returns
	```go
		func b() {
			for i := 0; i < 4; i++ {
				defer fmt.Println(i)
			}
		}
		* output: 3210
	```

	3. Deferred functions may read and assign to the returning function's named return value
	```go
		func c() (i int) {
			defer func() { i++ }()
			return 1
		}
		* output: 2
	```
*/

func main() {

}
