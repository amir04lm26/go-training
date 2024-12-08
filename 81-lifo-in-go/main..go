package main

import "fmt"

// NOTE: What is LIFO?
/*
	LIFO stands for Last In, First Out, which is a principle that describes how data is managed
	in certain data structures.
	It means that the last element added to a collection is the first one to be removed.

	In programming, LIFO is most commonly associated with stacks. A stack is a collection that
	operates on the LIFO principle, where:

	The push operation adds an element to the top of the stack (the end).
	The pop operation removes the top element from the stack (the last one added).
*/

// NOTE: LIFO in Go:
/*
	In Go, you can implement a stack using slices or linked lists. Here's a simple example using
	a slice:
*/

// NOTE: Explanation
/*
	In this example:
	We first push 10, 20, and 30 onto the stack.
	When we pop, we remove 30, the last element added.
	This LIFO behavior is useful in many scenarios, such as:

	Undo operations in text editors.
	Function call stacks in recursive algorithms.
	Depth-first search (DFS) in graph traversal.
*/

func main() {
	stack := []int{} // *  Initialize an empty stack

	// * Push elements to the stack
	// stack = append(stack, 10)
	// stack = append(stack, 20)
	// stack = append(stack, 30)
	stack = append(stack, 10, 20, 30)

	fmt.Println("Stack:", stack)

	// * Pop an element from the stack
	top := stack[len(stack)-1]   // * Get the last element (top)
	stack = stack[:len(stack)-1] // * Remove the last element

	fmt.Println("Popped:", top)
	fmt.Println("Updated Stack:", stack)
}
