package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

// NOTE: Sync.pool
/*
	sync.Pool in Go is a built-in type provided by the sync package that implements
	a pool of reusable objects. It is designed to minimize the overhead of allocating
	and deallocating objects, especially in high-performance scenarios where objects
	are frequently created and discarded.
*/

/// NOTE: Key Features of sync.Pool:
/*
	Object Reuse: Instead of creating a new object every time, sync.Pool allows you to
	reuse objects that have already been allocated, reducing memory allocation and
	garbage collection (GC) overhead.

	Thread-Safe: sync.Pool is safe to use across multiple goroutines concurrently.

	Lazy Allocation: If the pool is empty and no objects are available for reuse,
	it creates a new object on demand.

	GC Behavior: Objects in a sync.Pool may be garbage collected if they are not in
	use and no longer referenced, unlike traditional object pools where objects are
	permanently retained.
*/

// NOTE: How sync.Pool Works:
/*
	A sync.Pool manages a pool of objects with the following key methods:

	Get(): Retrieves an object from the pool. If the pool is empty, it allocates a new
	one by calling the New function.
	Put(x interface{}): Returns an object to the pool, making it available for reuse.
*/

// NOTE: Creating and Using a sync.Pool:

func syncPoolExample() {
	pool := sync.Pool{
		New: func() interface{} {
			// * Define how to create a new object
			fmt.Println("Creating new object")
			str := fmt.Sprintf("New Object %v", time.Now())
			return &str
		},
	}

	// * Get an object from the pool
	obj1 := pool.Get().(*string)
	fmt.Println("Got", *obj1)

	// * Put the object back into the pool
	pool.Put(obj1)

	// * Get the same object from the pool
	obj2 := pool.Get().(*string)
	fmt.Println("Got", *obj2)

	// * Pool is empty, creates a new object
	obj3 := pool.Get().(*string)
	fmt.Println("Got", *obj3)
}

// NOTE: Key Points to Note:
/*
	New Function: The New function is optional. If you don't provide it, and the pool is empty,
	Get() will return nil.

	Garbage Collection: Objects in the pool are not guaranteed to persist.
	The Go runtime can discard pooled objects during garbage collection, so you should not
	rely on sync.Pool for permanent storage.

	Scope of Use: sync.Pool is ideal for short-lived objects that are expensive to allocate,
	such as buffers or temporary structs used in high-frequency operations.

	Thread Safety: Multiple goroutines can safely call Get() and Put() on the same sync.Pool.
*/

// NOTE: Use Case Example: Buffer Pooling
/*
	Suppose you need a pool of byte buffers to handle temporary data.
*/

func bufferPoolingExample() {
	bufferPool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	fmt.Println("---------Second Example--------")

	// * Get a buffer from the pool
	buf := bufferPool.Get().(*bytes.Buffer)

	// *  Use the buffer
	buf.WriteString("Hello, sync.Pool!")
	fmt.Println("Buf", buf.String())

	// * Reset and put the buffer back into the pool
	buf.Reset()
	bufferPool.Put(buf)

	// * Get another buffer (might reuse the same one)
	newBuf := bufferPool.Get().(*bytes.Buffer)
	fmt.Println("Buf", newBuf.String()) // * Empty since it was reset
}

// NOTE: Advantages of sync.Pool:
/*
	Reduces memory allocations and garbage collection overhead.
	Improves performance for workloads requiring frequent, short-lived object creation.
	Limitations:
	Objects in the pool may still be garbage collected, so it's not a reliable cache.
	Should not be used for managing objects with long lifetimes or large, permanent data
	structures.
*/

// NOTE: Conclusion:
/*
	sync.Pool is a powerful tool for optimizing performance in scenarios where objects are
	frequently created and discarded. By reusing objects, it helps reduce allocation costs
	and GC pressure, making it particularly useful in high-concurrency applications like
	web servers or processing pipelines.
*/

func main() {
	syncPoolExample()
	bufferPoolingExample()
}
