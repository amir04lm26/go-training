package main

import (
	"fmt"
	"time"
)

// NOTE: Throttling Goroutines with Semaphore
/*
	Sometimes you need to limit the number of active goroutines in a more controlled way than a worker pool.
	A semaphore can help limit the number of concurrent goroutines without having a fixed worker pool size.
*/

// NOTE: Example: Using a Semaphore for Goroutine Throttling

// NOTE: Explanation:
/*
	•	Semaphore: The semaphore channel controls the number of active workers.
		Each worker sends an empty struct (struct{}{}) to the channel before starting, effectively acquiring a “permit”.
		When the semaphore is full (e.g., 2 permits), the remaining workers wait until a permit is released.
	•	Concurrency limit: In this example, only two workers are allowed to execute at the same time.
		The rest are blocked until one of the permits is released.

	In this example, only two workers can run simultaneously, controlled by the semaphore channel.
*/

func worker(id int, sem chan struct{}) {
	sem <- struct{}{} // * Acquire the semaphore
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d is done\n", id)
	<-sem // * Release the semaphore
}

func main() {
	semaphore := make(chan struct{}, 2) // * Limit to 2 concurrent workers

	for iterator := 1; iterator <= 5; iterator++ {
		go worker(iterator, semaphore)
	}

	time.Sleep(5 * time.Second)
}

// NOTE: What is semaphore?
/*
	A semaphore is a synchronization mechanism used in computer science and programming to control access to shared resources
	in a concurrent system, such as threads or processes. It helps prevent issues like race conditions, deadlocks, and resource
	contention by regulating how multiple entities access shared resources.

	Types of Semaphores
	1. Binary Semaphore:
		- Also known as a mutex (mutual exclusion lock).
		- Can only take two values: 0 (locked) or 1 (unlocked).
		- Ensures that only one thread or process can access the resource at any given time.
	2. Counting Semaphore:
		- Maintains a count of available resources.
		- Allows multiple threads or processes to access a limited number of instances of a resource.
		- Example: A counting semaphore initialized to 3 allows up to 3 threads to access the resource simultaneously.


	How a Semaphore Works
	1. Initialization:
		- A semaphore is initialized with a value, representing the number of available resources.
	2. Wait (P) Operation:
		- Decrements the semaphore's value.
		- If the value becomes less than zero, the process is blocked until the semaphore's value is incremented.
	3. Signal (V) Operation:
		- Increments the semaphore's value.
		- If any processes are waiting, one is unblocked and allowed to proceed.

	Examples in Real Life
	1. Printer Queue:
		- A semaphore ensures that only one job at a time can use a shared printer.
		- If a printer is busy, other print jobs are queued until it is free.
	2. Database Connection Pool:
		- A counting semaphore limits the number of simultaneous connections to a database.
		- Ensures the database isn't overwhelmed by too many connections.

	Code Example
	Binary Semaphore

	`import threading

	# Create a semaphore with an initial value of 1
	semaphore = threading.Semaphore(1)

	def critical_section(thread_id):
		print(f"Thread {thread_id} is waiting to access the resource.")
		semaphore.acquire()
		print(f"Thread {thread_id} has accessed the resource.")
		# Simulate resource usage
		import time
		time.sleep(2)
		print(f"Thread {thread_id} is releasing the resource.")
		semaphore.release()

	# Create and start multiple threads
	threads = []
	for i in range(5):
		t = threading.Thread(target=critical_section, args=(i,))
		threads.append(t)
		t.start()

	for t in threads:
		t.join()`

	Advantages of Semaphores
	1. Prevents Resource Conflicts: Ensures mutual exclusion for shared resources.
	2. Allows Controlled Concurrency: Useful in situations where limited access is needed.
	3. Simple Mechanism: Easy to implement in most programming languages.

	Disadvantages of Semaphores
	1. Complex to Manage: Mismanagement (e.g., failing to release) can lead to deadlocks.
	2. Not Scalable: As systems grow, semaphore-based synchronization can become inefficient.
	3. Potential for Starvation: Processes may wait indefinitely if not handled carefully.

	Semaphores are foundational tools in operating systems, threading libraries, and distributed systems,
	ensuring safe and efficient management of shared resources.
*/
