package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// NOTE: Synchronization Techniques
/*
	Synchronization techniques are methods used to coordinate and manage concurrent access to shared resources
	in multi-threaded or distributed systems, ensuring data consistency and preventing issues like race conditions,
	deadlocks, or resource contention. Below are commonly used synchronization techniques:
*/

// NOTE: 1. Locks and Mutexes
/*
	Locks: Prevent multiple threads or processes from accessing a shared resource simultaneously.
		- A thread acquires a lock before accessing the resource and releases it after it's done.

	Mutex (Mutual Exclusion):
		- A special type of binary semaphore.
		- Ensures only one thread accesses a resource at a time.

	Example:
	from threading import Lock

	lock = Lock()

	def critical_section():
		with lock:
			# Code that requires synchronization
			pass
*/

// NOTE: 2. Monitors
/*
	High-level synchronization constructs that combine mutual exclusion with condition variables.
	Used in object-oriented programming to encapsulate the resource, the condition for its access, and synchronization logic.
*/

// NOTE: 3. Semaphores
/*
	Binary Semaphore: Similar to a mutex, ensures mutual exclusion.
	Counting Semaphore: Allows a limited number of threads to access a resource simultaneously.
*/

// NOTE: 4. Barriers
/*
	- Synchronization points where threads or processes must wait until all have reached the barrier before continuing execution.
	- Often used in parallel computing.

	Example:
	from threading import Barrier

	barrier = Barrier(3)

	def task():
		print("Waiting at barrier")
		barrier.wait()
		print("Crossed the barrier")
*/

// NOTE: 5. Read-Write Locks
/*
	Specialized locks that allow:
		- Multiple readers to access a resource simultaneously (read-only).
		- Only one writer to modify the resource at a time.
	Improves performance for read-heavy workloads.
*/

// NOTE: 6. Condition Variables
/*
	- Used in conjunction with locks to enable threads to wait for certain conditions to be met.
	- Threads can release the lock and wait, then be notified when the condition is met.

	Example:
	from threading import Condition

	condition = Condition()

	def task():
		with condition:
			condition.wait()  # Wait for notification
			print("Condition met!")
*/

// NOTE: 7. Spinlocks
/*
	A lock where a thread continuously checks for availability instead of being put to sleep.
	Suitable for short wait times but can lead to CPU wastage.
*/

// NOTE: 8. Event Flags
/*
	- Use signaling mechanisms to indicate that an event has occurred.
	- Threads or processes can wait for these signals before proceeding.

	Example:
	from threading import Event

	event = Event()

	def task():
		event.wait()  # Wait for the event to be set
    	print("Event triggered!")
*/

// NOTE: 9. Message Passing
/*
	- Avoids shared memory by sending messages between threads or processes.
	- Ensures synchronization implicitly as messages are processed in order.

	Example: Using Python's queue module:
	from queue import Queue

	queue = Queue()

	def producer():
		queue.put("Data")

	def consumer():
		data = queue.get()
		print(f"Consumed {data}")
*/

// NOTE: 10. Atomic Operations
/*
	Operations that are indivisible and automatically synchronized at the hardware level.
	Examples: Atomic counters, compare-and-swap (CAS).
*/

// NOTE: 11. Transactions
/*
	Ensures that a group of operations is executed as a single, atomic unit (all or nothing).
	Common in databases with ACID (Atomicity, Consistency, Isolation, Durability) properties.
*/

// NOTE: 11. Transactions
/*
	Ensures that a group of operations is executed as a single, atomic unit (all or nothing).
	Common in databases with ACID (Atomicity, Consistency, Isolation, Durability) properties.
*/

// NOTE: 12. Distributed Synchronization
/*
	Techniques for distributed systems:
		- Consensus Algorithms (e.g., Paxos, Raft): Ensure agreement among nodes.
		- Leases: Temporary locks in distributed systems.
		- Vector Clocks: For tracking causality.
*/

// NOTE: 13. Software Transactional Memory (STM)
/*
	Allows threads to execute in parallel while ensuring that any shared memory operations are atomic and consistent.
	Used in functional programming languages like Haskell.
*/

// NOTE: 14. Priority Inheritance
/*
	A technique to prevent priority inversion by temporarily increasing the priority of a thread holding a resource
	needed by a higher-priority thread.
*/

// NOTE: Choosing the Right Synchronization Technique
/*
	The choice depends on:
	- Concurrency Model: Thread-based, process-based, or event-driven.
	- Resource Type: Read-heavy, write-heavy, or mixed.
	- Performance Requirements: Low-latency, high-throughput, or fault tolerance.
	- System Type: Single system or distributed.

	Each technique has its trade-offs, and often, multiple techniques are combined for optimal performance and correctness.
*/

// * NOTE: Synchronization in Go
// NOTE: 1. Locks and Mutexes
/*
	Locks and mutexes are the most basic synchronization primitives. They ensure that only one thread or
	goroutine can access a shared resource at a time.

	Details
		- A lock allows a thread/goroutine to "lock" a resource while it is being used and "unlock" it when done.
		- A mutex (mutual exclusion) is a type of lock that prevents multiple goroutines from entering a critical
		  section simultaneously.
*/

// NOTE: Example in Go

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	fmt.Println("Counter:", counter)
	mutex.Unlock()
}

func lockAndMutexExample() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
}

// NOTE: 2. Monitors
/*
	Monitors combine mutual exclusion with condition variables, providing a higher-level abstraction for synchronization.

	Details
		- Go does not directly have a monitor construct but provides sync.Cond, which implements condition
		  variables for signaling among goroutines.
*/

// NOTE: Example in Go

var condition = sync.NewCond(&sync.Mutex{})

func monitorWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	condition.L.Lock()
	condition.Wait() // * Wait for a signal
	fmt.Printf("Worker %d started\n", id)
	condition.L.Unlock()
}

func monitorExample() {
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go monitorWorker(iterator, &wg)
	}

	fmt.Println("Progress started...")
	time.Sleep(time.Second)
	condition.L.Lock()
	fmt.Println("Signaling workers...")
	condition.Broadcast() // * Notify all waiting goroutines
	condition.L.Unlock()

	wg.Wait()
}

// NOTE: 3. Semaphores
/*
	Semaphores control access to a resource by maintaining a count. Go doesn’t have a built-in semaphore,
	but you can implement one using channels.

	Details
		- Binary Semaphore: Allows one goroutine at a time.
		- Counting Semaphore: Allows a limited number of goroutines to access a resource.
*/

// NOTE: Example in go
func semaphoreWorker(id int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	semaphore <- struct{}{} // *  Acquire semaphore
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
	<-semaphore
}

func semaphoreExample() {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 2) // * Allow 2 workers at a time

	for iterator := 1; iterator <= 5; iterator++ {
		wg.Add(1)
		go semaphoreWorker(iterator, semaphore, &wg)
	}

	wg.Wait()
}

// NOTE: 4. Barriers
/*
	Barriers force all goroutines to reach a specific point before continuing execution.

	Details
		- Go does not have built-in barriers, but you can achieve this using sync.WaitGroup or custom logic.
*/

// NOTE: Example in Go
func barrierWorker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d waiting at the barrier\n", id)
	barrier.Wait()
	fmt.Printf("Worker %d passed the barrier\n", id)
}

func barrierExample() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	// * Set up the barrier
	barrier.Add(1)

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go barrierWorker(iterator, &wg, &barrier)
	}

	// * Simulate reaching the barrier
	time.Sleep(time.Second)
	fmt.Println("All workers have reached the barrier")
	barrier.Done() // * Release all workers
	wg.Wait()
}

// NOTE: 5. Read-Write Locks
/*
	Allows multiple readers but only one writer. Useful for scenarios where reads are more frequent than writes.

	Details
		- Use sync.RWMutex in Go for read-write locks.
*/

// NOTE: Example in Go

var (
	rwCounter int
	rwMutex   sync.RWMutex
)

func rwMutexReader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Reader %d read counter: %d\n", id, rwCounter)
	rwMutex.RUnlock()
}

func rwMutexWriter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock()
	rwCounter++
	fmt.Printf("Writer updated counter to: %d\n", rwCounter)
	rwMutex.Unlock()
}

func rwMutexExample() {
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 3; iterator++ {
		wg.Add(1)
		go rwMutexReader(iterator, &wg)
	}

	wg.Add(1)
	go rwMutexWriter(&wg)

	wg.Wait()
}

// NOTE: 6. Condition Variables
/*
	Condition variables allow threads or goroutines to wait for a condition and get notified when it is met.

	Details
		- Use sync.Cond in Go.
*/

// NOTE: Example in Go

var (
	cond     = sync.NewCond(&sync.Mutex{})
	resource = false
)

func condConsumer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for !resource {
		cond.Wait() // * Wait for the resource to be available
	}
	fmt.Printf("Consumer %d consumed the resource\n", id)
	cond.L.Unlock()
}

func condProducer() {
	cond.L.Lock()
	resource = true
	fmt.Println("Producer created the resource")
	cond.Signal() //*  Notify one waiting goroutine
	cond.L.Unlock()
}

func condExample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go condConsumer(1, &wg)

	// * Simulate some work
	condProducer()

	wg.Wait()
}

// NOTE: 7. Spinlocks
/*
	Spinlocks continuously check for resource availability, often implemented with a busy-wait loop.
	Go doesn’t have a built-in spinlock, but it can be simulated.
*/

// NOTE: 8. Message Passing
/*
	Message passing avoids shared memory by using channels or queues. This is idiomatic in Go.
*/

// NOTE: Example in Go
func messageProducer(ch chan int) {
	for iterator := 0; iterator < 5; iterator++ {
		ch <- iterator
		fmt.Println("Produced:", iterator)
	}
	close(ch)
}

func messageConsumer(ch chan int) {
	for item := range ch {
		fmt.Println("Consumed:", item)
	}
}

func messageExample() {
	ch := make(chan int, 2)

	go messageProducer(ch)
	messageConsumer(ch)
}

// NOTE: 9. Atomic Operations
/*
	Atomic operations ensure that read-modify-write operations on shared resources are performed
	as a single, indivisible step. This prevents race conditions without the need for locks.

	Details
		- Atomic operations are fast and avoid the overhead of traditional locking mechanisms.
		- In Go, the sync/atomic package provides atomic operations like AddInt32, LoadInt32,
		  and CompareAndSwapInt32.
*/

// NOTE: Example in Go
var atomicCounter int32

func atomicIncrement(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&atomicCounter, 1)
}

func atomicExample() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go atomicIncrement(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", atomicCounter)
}

// NOTE: 10. Transactions
/*
	Transactions are used to ensure a group of operations is performed atomically and consistently.
	They are common in databases but can also be implemented in synchronization contexts using
	locks or other mechanisms.

	Details
		- Transactions follow the ACID properties (Atomicity, Consistency, Isolation, Durability).
		- In Go, explicit transaction-like behavior is implemented manually using locks.
*/

// NOTE: Example in Go
type Account struct {
	balance int
	mu      sync.Mutex
}

func (account *Account) transfer(amount int, target *Account) {
	account.mu.Lock()
	defer account.mu.Unlock()

	target.mu.Lock()
	defer target.mu.Unlock()

	if account.balance > amount {
		account.balance -= amount
		target.balance += amount
		fmt.Printf("Transferred %d, Remaining Balance: %d\n", amount, account.balance)
	} else {
		fmt.Println("Insufficient amount for transfer")
	}
}

func transactionExample() {
	account1 := &Account{balance: 100}
	account2 := &Account{balance: 50}

	account1.transfer(20, account2)

	fmt.Println("Account 1 Balance:", account1.balance)
	fmt.Println("Account 2 Balance:", account2.balance)
}

// NOTE: 11. Spinlocks
/*
	A spinlock uses a busy-wait loop to acquire a lock. It is useful for very short
	critical sections but can waste CPU cycles if contention is high.

	Details
		- Go doesn’t have built-in spinlocks, but they can be implemented using atomic operations.
*/

// NOTE: Example in Go

type Spinlock struct {
	flag int32
}

func (spinlock *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(&spinlock.flag, 0, 1) {
		// * Busy-wait
	}
}

func (spinlock *Spinlock) Unlock() {
	atomic.StoreInt32(&spinlock.flag, 0)
}

func spinLockExample() {
	var spin Spinlock

	spin.Lock()
	fmt.Println("Acquired spinlock")

	go func() {
		spin.Lock()
		fmt.Println("Goroutine acquired spinlock")
		spin.Unlock()
	}()

	time.Sleep(time.Second)
	spin.Unlock()
	fmt.Println("Released spinlock")
}

// NOTE: 12. Readers-Writer Problem
/*
	Latches ensure that one or more goroutines wait until a specific condition is met,
	often implemented with sync.WaitGroup.

	Details
		- Latches can block goroutines until certain operations complete.
*/

// NOTE: Example in Go

func latchWorker(id int, latch *sync.WaitGroup) {
	defer latch.Done()
	fmt.Printf("Worker %d completed\n", id)
}

func latchExample() {
	var latch sync.WaitGroup

	for iterator := 1; iterator < 5; iterator++ {
		latch.Add(1)
		go latchWorker(iterator, &latch)
	}

	latch.Wait() // * Wait for all workers to finish
	fmt.Println("All workers are done")
}

// NOTE: 15. Event Loops
/*
	Event loops process events or tasks in a loop, ensuring tasks are executed sequentially
	while avoiding locks by using queues.

	Details
		- Go uses channels and select statements to implement event loops.
*/

func eventLoop(events <-chan string) {
	for event := range events {
		fmt.Println("Processing:", event)
	}
}

func eventLoopExample() {
	events := make(chan string, 5)

	go eventLoop(events)

	events <- "Event 1"
	events <- "Event 2"
	events <- "Event 3"

	time.Sleep(time.Second)
	close(events)
}

/*
Each synchronization technique has unique use cases depending on the type of concurrency
problem you're solving.
*/
func main() {
	lockAndMutexExample()
	monitorExample()
	semaphoreExample()
	barrierExample()
	rwMutexExample()
	condExample()
	messageExample()
	atomicExample()
	transactionExample()
	spinLockExample()
	latchExample()
	eventLoopExample()
}
