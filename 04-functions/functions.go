package main

import "fmt"

// NOTE: When we define function, it gets parameters
// NOTE: When we use the function, it gets arguments

// NOTE: GC Optimization
// NOTE: 1. Function Inlining
/*
	Golang’s compiler inlines small, simple functions automatically to improve performance
	by avoiding the overhead of function calls. This is why sometimes Go “uses the
	functionality inside the function rather than calling the function”.

	Example:
	```
		package main

		import "fmt"

		func add(a, b int) int { // Small function
			return a + b
		}

		func main() {
			x := add(5, 10) // This may be inlined
			fmt.Println(x)
		}
	```

	How it works:
	•	Instead of actually calling add(5, 10), the compiler may replace it with x := 5 + 10,
		eliminating the function call entirely.

	When does Go inline a function?
	•	The function is small and simple.
	•	No complex loops or recursion inside the function.
	•	No defer statements inside the function.
	•	The function does not have go (goroutine) or recover.

	Checking if a function is inlined:
	You can use the go build -gcflags="-m" command:
	```
		go build -gcflags="-m" main.go
	```
	It will output
	```
		./main.go:6:6: can inline add
	```
	indicating that add() is inlined.
*/

// NOTE: 2. Escape Analysis & Stack Allocation
/*
	Another optimization in Go is escape analysis, where the compiler decides whether to
	allocate variables on the stack or heap.
	•	If a variable does not escape a function, it is allocated on the stack (faster).
	•	If a variable escapes the function (e.g., returned as a pointer), it is allocated on the heap (slower).

	Example:
	```
		func createNumber() *int {
			num := 42
			return &num  // Escapes to heap
		}
	```

	Here, num is allocated on the heap because it’s returned as a pointer.

	To check escape analysis:
	```
		go build -gcflags="-m" main.go
	```

	Output: ./main.go:4:9: &num escapes to heap
	```
		./main.go:4:9: &num escapes to heap
	```
*/

// NOTE: 3. Inlining vs. Function Call Overhead
/*
	Inlining is beneficial when:
		•	Function call overhead is significant (e.g., in loops).
		•	The function is simple (e.g., getter functions, arithmetic operations).

	However, excessive inlining can increase binary size. The Go compiler decides
	inlining automatically to balance performance and binary size.
*/

// NOTE: 4. Loop Unrolling & Bound Check Elimination
/*
	•	Go’s compiler optimizes loops, sometimes unrolling them (expanding the loop
		body to avoid repeated checks).
	•	It can also remove array bounds checks if it can prove they are unnecessary.
*/

// NOTE: Summary:
/*
	•	Inlining replaces function calls with function bodies to improve performance.
	•	Escape analysis decides whether variables should be on the stack or heap.
	•	Loop optimizations reduce overhead.
	•	Use go build -gcflags="-m" to check inlining and escape analysis.
*/

// NOTE: When a variable escapes to the heap?
/*
	A variable escapes to the heap when the Go compiler determines that it cannot be safely
	allocated on the stack. This typically happens in the following scenarios:
*/

// NOTE: 1. Returning a Pointer to a Local Variable
/*
	If a function returns a pointer to a local variable, the variable must be stored on the
	heap because the stack frame disappears after the function returns.

	Example:
	```
		package main

		import "fmt"

		func createPointer() *int {
			num := 42 // num is a local variable
			return &num // num escapes to the heap
		}

		func main() {
			p := createPointer()
			fmt.Println(*p) // Still accessible because it's on the heap
		}
	```

	Why?
		•	num is allocated inside createPointer(), but since we return a pointer to num,
			it must survive beyond the function call.
		•	The stack frame of createPointer() is destroyed after execution, so num moves
			to the heap to remain accessible.

	Checking with escape analysis:
	```
		go build -gcflags="-m" main.go
	```

	Output:
	```
		./main.go:6:9: &num escapes to heap
	```
*/

// NOTE: 2. Storing a Variable in an Interface
/*
	If a local variable is assigned to an interface, it escapes to the heap because
	interfaces in Go store values indirectly.

	Example:
	```
		package main

		import "fmt"

		func main() {
			var x int = 10
			var i interface{} = x // x escapes to heap
			fmt.Println(i)
		}
	```

	Why?
	•	interface{} is a dynamically-typed construct.
	•	Go must allocate x on the heap to ensure its value remains accessible.
*/

// NOTE: 3. Appending to a Slice Beyond Capacity
/*
	If a slice grows beyond its initial capacity, Go reallocates it on the heap.

	Example:
	```
		package main

		func main() {
			s := make([]int, 2, 2) // Stack allocation (capacity 2)
			s = append(s, 3)       // New allocation on heap
		}
	```

	Why?
	•	When append() exceeds capacity, Go creates a new slice on the heap and
		copies the data over.
*/

// NOTE: 4. Using new() or make()
/*
	Allocating objects explicitly using new() or make() often results in heap allocation.

	Example:
	```
		package main

		func main() {
			p := new(int) // Always allocated on heap
			*p = 42
		}
	```

	Why?
	•	new(int) explicitly requests heap allocation


	new() is a built-in function that allocates memory for a zeroed value of a specified type and returns a pointer to it.

	Explanation:
	Type can be any Go type (e.g., int, struct, float64, etc.).
	new(Type) returns a pointer of type *Type.
	The value is initialized to the zero value for that type (e.g., 0 for numbers, "" for strings,
	nil for slices/maps/interfaces/pointers, etc.).

	Key points:
	new(Type) may allocate memory on the heap if the compiler determines that the pointer "escapes"
	the current function (i.e., is used outside its scope).

	If the pointer does not escape, the compiler may allocate it on the stack, even if you use new().
*/

// NOTE: 5. Closures Capturing Variables
/*
	If a function captures a local variable in a closure, the variable escapes to the heap.

	Example:
	```
		package main

		import "fmt"

		func main() {
			var num = 42
			f := func() { fmt.Println(num) } // num escapes to heap
			f()
		}
	```

	Why?
	•	The anonymous function (f) captures num, so num must outlive main()’s stack frame.
	•	To make this possible, Go moves num to the heap.
*/

// NOTE: 6. Structs Passed to Channels or Goroutines
/*
	If a variable is passed to a goroutine or a channel, it escapes because
	the Go scheduler cannot guarantee that the stack frame will persist.

	Example:
	```
		package main

		func main() {
			msg := "hello"
			go func() {
				println(msg) // msg escapes to heap
			}()
		}
	```

	Why?
	•	The goroutine might run after main() has returned.
	•	msg is stored on the heap to keep it accessible.
*/

// NOTE: How to Prevent Heap Escapes?
/*
	1.	Use value types instead of pointers whenever possible:
	```
		func createValue() int {
			num := 42
			return num // num stays on stack
		}
	```
	2.	Pass small structs by value instead of pointers.
	3.	Use stack-allocated slices by controlling capacity:
	```
		s := make([]int, 5) // Preallocated, stays on stack
	```
	4.	Avoid unnecessary interfaces if exact types can be used.
*/

// NOTE: Summary
/*
	What causes a variable to escape to the heap?

	✅ Returning a pointer to a local variable
	✅ Assigning a variable to an interface
	✅ Growing a slice beyond capacity
	✅ Using new() or make()
	✅ Capturing variables in a closure
	✅ Passing variables to goroutines or channels
*/

// NOTE: Want to check your code? Use:
/*
	```
		go build -gcflags="-m" main.go
	```
*/

// NOTE: Example of Bound Check Elimination in Go
/*
	Bound check elimination is an optimization where the Go compiler removes unnecessary
	array/slice bounds checking to improve performance.
*/

// NOTE: 1. What is Bound Check?
/*
	In Go, accessing an array or slice element out of bounds will cause a runtime panic. To prevent this, Go automatically inserts bounds checks for each access.

	Example with bounds checking:
	```
		func getElement(arr []int, index int) int {
			return arr[index] // Compiler inserts a bounds check
		}
	```
	If index >= len(arr), Go will panic to prevent unsafe memory access.
*/

// NOTE: 2. How Does Go Eliminate Bound Checks?
/*\
Go removes bounds checks when it can prove at compile time that the index is within range.

✅ Example of Bound Check Elimination
```
	package main

	import "fmt"

	func sumArray(arr []int) int {
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum += arr[i] // Bounds check eliminated
		}
		return sum
	}

	func main() {
		arr := []int{1, 2, 3, 4, 5}
		fmt.Println(sumArray(arr))
	}
```

Why is the bounds check eliminated?
•	The loop condition ensures that i is always < len(arr), so an out-of-bounds access is impossible.
•	The compiler removes redundant bounds checks because it can mathematically prove the access is safe.
*/

// NOTE: Checking Bounds Check Elimination
/*
	Run:
	```
		go build -gcflags="-d=ssa/check_bce/debug=1" main.go
	```

	If successful, you will see:
	```
		./main.go:9:12: Found IsInBounds
		./main.go:9:12: Removed IsInBounds
	```
	This means the bounds check was removed.
*/

// NOTE: 3. When Is Bound Check NOT Eliminated?
/*
	❌ Example where bounds check remains:
	```
		func unsafeAccess(arr []int, i int) int {
			return arr[i] // Compiler cannot guarantee i is in range
		}
	```

	Why?
	•	i is an unknown value at runtime, so Go must keep the bounds check.
*/

// NOTE: 4. Ways to Ensure Bound Check Elimination
/*
	✅ Use range Loops
	```
		func sumRange(arr []int) int {
			sum := 0
			for _, v := range arr {
				sum += v // No bounds check needed
			}
			return sum
		}
	```
	•	range ensures safe iteration without explicit indexing.

	✅ Manually Hoist Bounds Check
	```
		func sumHoisted(arr []int) int {
			sum := 0
			n := len(arr) // Hoist length check
			for i := 0; i < n; i++ {
				sum += arr[i] // Bounds check eliminated
			}
			return sum
		}
	```

	•	Storing len(arr) in a variable helps Go prove bounds safety.
*/

// NOTE: 5. Summary
/*
	Code Pattern						Bounds Check Eliminated?		 Reason
	for i := 0; i < len(arr); i++		✅ Yes							Compiler sees i is safe
	for _, v := range arr				✅ Yes							No index access needed
	Using dynamic index (arr[i])		❌ No							i is unknown at compile time

	To check bounds elimination, use:
	```
		go build -gcflags="-d=ssa/check_bce/debug=1" main.go
	```
*/

func greet(name string) {
	println("Hello, ", name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}

func main() {
	greet("Amir")

	added := add(23, 62)
	println("Some up value:", added)

	divided, err := divide(10, 2)
	if err != nil {
		println("Error:", err)
	} else {
		println("divided:", divided)
	}

	divideRes, divideErr := divide(10, 0)
	if divideErr != nil {
		fmt.Println("Error:", divideErr)
	} else {
		println("divided:", divideRes)
	}

	// * called anonymous func in go
	func() {
		value := "IIFE"
		fmt.Printf("I believe this is an immediately invoked function expression (%v)\r\n", value)
	}()
	// value = "error" > error > undefined
}
