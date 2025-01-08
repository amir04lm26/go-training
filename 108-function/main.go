package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type personWithMethod struct {
	name string
	age  int
}

// * methods

func (person *personWithMethod) greeting() string {
	return fmt.Sprintf("Hello, I'm %v", person.name)
}

// * interfaces

type employee struct {
	person personWithMethod
	id     int
}

func (person *employee) greeting() string {
	return fmt.Sprintf("Hello employees, I'm %v", person.person.name)
}

// NOTE: An interface in go defines set of method signatures
// NOTE: Polymorphism is the ability of a VALUE of a certain TYPE to be also another TYPE
type human interface {
	greeting() string
}

func humanGreeting(human human) {
	fmt.Println(human.greeting())
}

// * stringer interface
type count int

func (count count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(count)))
}

// NOTE: log.Println vs fmt.Println
/*
	log.Println has timestamp attached to it. Also, we can use something else rather than stdout
*/

// * wrapper
// NOTE: wrapper is a function that provides an additional layer of abstraction or functionality around an existing function
func logInfo(stringer fmt.Stringer) {
	log.Println("LOG FROM LOGGER", stringer.String())
}

func main() {
	// NOTE: func (receiver) identifier(parameters) (returns) { code }
	// * func (receiver) identifier(parameter(s)) (return(s)) { code } (argument(s))

	result := sum(2, 3, 5, 6)
	fmt.Println("result:", result)

	result2 := sum() // * passing nil
	fmt.Println("result2:", result2)

	// NOTE: variadic parameter unpacking or slice unpacking
	inputs := []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9}
	result3 := sum(inputs...)
	fmt.Println("result3:", result3)

	// * Running immediately after the surrendering function returns
	// * either executing return statement, reaching the end of function body or because the compounding goroutine is panicking
	defer deferred()

	time.Sleep(time.Second)

	p1 := personWithMethod{"Amir", 27}
	fmt.Println(p1.greeting())

	p2 := employee{
		person: personWithMethod{"Faeze", 24},
		id:     123,
	}
	fmt.Println(p2.greeting())

	// * interface
	humanGreeting(&p1)
	humanGreeting(&p2)

	// * stringer interface
	var num count = 20
	fmt.Println(num)

	// * log wrapper
	log.Println(num)
	logInfo(num)

	// * anonymous function - encapsulate and limit the scope
	// NOTE: func (parameters) (returns) { code }
	func() {
		fmt.Println("Anonymous func running")
	}()

	res := func(str string) string {
		fmt.Println(str)
		return str
	}("anonymous func in go")
	fmt.Println("res:", res)

	// * function expression
	fn := func() {
		fmt.Println("This is function expression")
	}
	fn()
	fmt.Printf("fn type is: %T\n", fn)

	// NOTE: Expression
	/*
		An expression is a combination of values, variables, operators and function calls
		that are evaluated to produce a value
	*/

	// NOTE: a statement is a combination of expressions

	// * return a function
	fmt.Println("withReturn:", withReturn()())

	// * closure
	closure := withReturn()
	fmt.Println("closure result:", closure())
	fmt.Println("closure result:", closure())
	fmt.Println("closure result:", closure())
	fmt.Println("closure result:", closure())

	// * callback
	fmt.Println("withCallback:", withCallback(sum))

	// * Recursion
	// * A recursive function is a function that calls itself during its execution until it reaches
	// * a base case, which is a simple case that can be solved directly without further recursion.
	fmt.Println("4! is:", factorial(4))
}

func sum(values ...int) (result int) {
	for _, num := range values {
		// fmt.Println(num)
		result += num
	}

	return
}

func deferred() {
	fmt.Println("Now the deferred function is executing")
}

// * return a function
func withReturn() func() int {
	// * closure
	num := 42

	return func() int {
		num++
		return num
	}
}

// * callback

// func withCallback(callback func(...int) int) int {
// 	return callback(42, 43, 44)
// }

func withCallback(callback func(num ...int) int) int {
	return callback(42, 43, 44)
}

// * recursion

func factorial(num int) int {
	if num < 0 {
		log.Fatal("Invalid input")
	}
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
