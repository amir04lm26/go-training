package main

import "fmt"

// NOTE: Stack and Heap
/*
	Stack and Heap are two types of memory used by programs to store data during execution.
	They differ in terms of how they allocate, manage, and release memory.
*/

// NOTE: Stack
/*
	The stack is a region of memory that stores data in a last-in, first-out (LIFO) manner.
	It is primarily used for function calls, local variables, and control flow.

	Characteristics of the Stack
	1. Fast Access:
		Memory allocation and deallocation on the stack are very fast because it uses a
		simple pointer to manage memory.

	2. Automatic Management:
		Memory is automatically allocated when a function is called and deallocated when
		the function returns.

	3. Local Scope:
		The stack is used to store local variables and function call information like
		parameters, return addresses, and temporary data.

	4. Limited Size:
		The stack is relatively small, and excessive use (e.g., infinite recursion) can
		cause a stack overflow.

	5. Fixed Lifetime:
		Data on the stack exists only as long as the function or block in which it is
		declared is active.
*/

// NOTE: Example of Stack Memory in Go:
/*
	func add(a int, b int) int {
		result := a + b // `result` is stored on the stack
		return result   // Memory is freed when the function exits
	}
*/

// NOTE: Heap
/*
	The heap is a region of memory used for dynamic memory allocation.
	It is managed manually or automatically through mechanisms like garbage collection.

	Characteristics of the Heap
	1. Dynamic Allocation:
		The heap is used to allocate memory at runtime, typically for data structures like
		slices, maps, and pointers.

	2. Global Scope:
		Data on the heap is accessible as long as a reference to it exists, even after the
		function that allocated it exits.

	3. Managed Lifetime:
		Memory allocated on the heap must be explicitly freed in some languages (like C),
		but in Go, garbage collection automatically cleans up unused memory.

	4. Slower Access:
		Allocating and accessing memory on the heap is slower than the stack due to the
		need for more complex memory management.

	5. Large Size:
		The heap has a much larger size limit compared to the stack.
*/

// NOTE: Example of Heap Memory in Go:
/*
	func createPointer() *int {
		value := 42        // Local variable `value` on the stack
		return &value      // Address of `value` is returned, stored on the heap
	}

	func main() {
		p := createPointer() // `p` points to memory in the heap
		fmt.Println(*p)      // Accessing heap memory
	}

	In this example, the variable value escapes to the heap because its pointer is returned
	and used outside the function scope.
*/

// NOTE: Key Differences Between Stack and Heap
/*
	Aspect				Stack								Heap
	Storage Method		LIFO (Last In, First Out)			Random access
	Speed				Faster								Slower
	Lifetime of Data	Tied to function/block scope		Exists as long as a reference exists
	Management			Automatically managed				Garbage collected in Go
	Typical Use Case	Local variables, function calls		Dynamically allocated memory
	Size				Smaller								Larger
*/

// NOTE: Why Does This Matter in Go?
/*
	1. Escape Analysis:
		Go's compiler determines whether a variable should be allocated on the stack or heap
		based on its escape analysis. If a variable's lifetime exceeds the scope of its
		function (e.g., it’s returned as a pointer), it is moved to the heap.

	2. Efficiency:
		Understanding stack vs. heap helps write efficient code by minimizing unnecessary heap
		allocations.

	3. Garbage Collection:
		Go manages heap memory using garbage collection, which simplifies memory management
		but can introduce some performance overhead.
*/

// NOTE: Example Demonstrating Both Stack and Heap in Go
/*
	The x variable in stackExample is allocated and freed on the stack.
	The Person struct in heapExample is allocated on the heap because its pointer escapes
	the function scope.
*/

// NOTE: Summary
/*
	Stack: Faster, smaller, and used for local variables and function calls.
	Heap: Larger, slower, and used for dynamic or long-lived memory.
	In Go, garbage collection simplifies heap memory management, while the compiler's
	escape analysis determines whether variables should go on the stack or heap.
*/

type Person struct {
	Name string
}

func stackExample() {
	x := 10 // Stored on the stack
	fmt.Println(x)
}

func heapExample() *Person {
	p := &Person{Name: "Alice"} // Stored on the heap
	return p                    // Reference to heap memory is returned
}

func main() {
	stackExample()
	person := heapExample()
	fmt.Println(person.Name)
}
