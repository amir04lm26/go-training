package main

import (
	"fmt"
	"sync"
)

// NOTE: RWMutex (Read-Write Mutex)
/*
	An RWMutex is a specialized mutex that allows multiple readers or a single writer to access
	shared data.
	It provides better performance than a regular mutex when you have many concurrent reads and
	fewer writes.
*/

// NOTE: Example: Using RWMutex for Read-Write Locking

// NOTE: Explanation:
/*
	•	sync.RWMutex: The RWMutex provides separate methods for read locking (RLock) and write
		locking (Lock).
	•	Get method: This method uses a read lock (RLock) to allow multiple readers to access
		the map concurrently.
	•	Set method: This method uses a write lock (Lock) to ensure exclusive access when
		modifying the map.

	The use of RWMutex allows for efficient concurrent access to the map, with multiple readers
	able to access it concurrently while writes are serialized.
*/

type SafeMap struct {
	rwMu     sync.RWMutex
	valueMap map[string]int
}

func (safeMap *SafeMap) Get(key string) int {
	safeMap.rwMu.RLock()
	defer safeMap.rwMu.RUnlock()
	return safeMap.valueMap[key]
}

func (safeMap *SafeMap) Set(key string, value int) {
	safeMap.rwMu.Lock()
	defer safeMap.rwMu.Unlock()
	safeMap.valueMap[key] = value
}

func main() {
	/*
		The make function in Go is used to create and initialize certain types of objects, specifically
		slices, maps, and channels. In the snippet you provided:

		The make(map[string]int) part initializes an empty map with keys of type string and values of
		type int. This ensures that the map is ready to use, and you can safely add, retrieve, or
		delete key-value pairs from it.
	*/
	safeMap := SafeMap{valueMap: make(map[string]int)}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Set("key1", 100)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Key1:", safeMap.Get("key1"))
	}()

	wg.Wait()
}
