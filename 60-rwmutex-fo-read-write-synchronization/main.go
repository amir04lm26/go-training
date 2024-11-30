package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: RWMutex for Read-Write Synchronization
/*
	In cases where multiple goroutines need to read shared data concurrently but write access is exclusive, Go provides sync.RWMutex.
	It allows multiple readers or a single writer at any given time.
*/

// NOTE: Explanation:
/*
	•	RWMutex: sync.RWMutex provides two types of locks: RLock for readers and Lock for writers.
		Multiple readers can hold the lock simultaneously, but writers need exclusive access.
	•	RLock/Lock: Readers use RLock to acquire the lock for reading, and writers use Lock to acquire it for writing.
*/

// NOTE: Output
/*
	Reader 6: reading data = 0
	Write 2: writing data = 2
	Reader 4: reading data = 2
	Reader 2: reading data = 2
	Reader 5: reading data = 2
	Reader 1: reading data = 2
	Reader 3: reading data = 2
	Write 1: writing data = 1

	The readers can access the shared data concurrently, while the writers need exclusive access.
	In this example, the readers print the data after each writer updates it.
*/

// NOTE: Example: Using RWMutex for Read-Write Locking
var data int
var rwMu sync.RWMutex

func readData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.RLock() // * Acquire read lock
	fmt.Printf("Reader %d: reading data = %d\n", id, data)
	time.Sleep(time.Second)
	rwMu.RUnlock() // * Read read lock
}

func writeData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.Lock() // * Acquire write lock
	fmt.Printf("Write %d: writing data = %d\n", id, id)
	data = id
	time.Sleep(time.Second)
	rwMu.Unlock() // * Release write lock
}

func main() {
	var wg sync.WaitGroup

	for iterator := 1; iterator <= 2; iterator++ {
		wg.Add(1)
		go writeData(iterator, &wg)
	}

	for iterator := 1; iterator <= 6; iterator++ {
		wg.Add(1)
		go readData(iterator, &wg)
	}

	wg.Wait()
}
