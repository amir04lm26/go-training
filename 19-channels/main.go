package main

import "fmt"

// NOTE
/*
	1. Channel (chan):

		•	Channels in Go are a way to send and receive data between Goroutines.
		•	In the code, resultChannel is a channel that allows integers (chan int) to be passed between Goroutines.
		•	You use the <- operator to send data to or receive data from a channel.
		•	For example, resultChannel <- sum sends the sum value to the channel, and firstHalfSum := <- resultChannel receives a value from the channel.

	2. Goroutines (go keyword):

		•	A Goroutine is a lightweight thread managed by Go’s runtime. They are created using the go keyword.
		•	In this example, two Goroutines are launched:
		•	One sums the first half of the array: go sumOfSlice(array[:len(array)/2], resultChannel).
		•	The other sums the second half: go sumOfSlice(array[len(array)/2:], resultChannel).

	3. Range (range keyword):

		•	range is used to iterate over elements in various data structures like arrays, slices, maps, or channels.
		•	In this example, for _, value := range numbers is used to loop through each element of the numbers slice.
		•	_ means the index of the element is ignored.
		•	value is the element itself.
		•	range allows for clean iteration over slices without needing manual indexing.

	4. Array Slicing ([:]):

		•	array[:len(array)/2] slices the original array to get the first half.
		•	array[len(array)/2:] slices it to get the second half.
		•	This is how the array is split into two parts for parallel processing.

	5. Channel Blocking and Synchronization:

		•	Channels are blocking. When you try to receive a value from a channel (e.g., firstHalfSum := <- resultChannel), the Goroutine will pause until another Goroutine sends a value to the channel.
		•	In this example, x, y := <-resultChannel, <-resultChannel ensures that the main Goroutine waits for both results from the two Goroutines before printing the sum.

	6. Summing Values:

		•	After both Goroutines send their partial sums to the channel, the main Goroutine receives those values (firstHalfSum and secondHalfSum) and adds them together to get the total sum.

*/

// NOTE: Close channel
/*
	The close function is used to close a channel.
	Once a channel is closed, no more values can be sent to it, but it can still be read from.
	It’s good practice to close a channel when you are done sending values to
	avoid a Goroutine blocking indefinitely when trying to read from an empty channel.
*/

// NOTE: What Happens When a Channel is Closed?
/*
	•	Sending on a closed channel causes a panic.
	•	Receiving from a closed channel returns the remaining values in the buffer.
		Once the buffer is empty, any further receives will return
		the zero value of the channel’s type and will not block.
*/

/*
	Recap:

	•	chan is used for communication between Goroutines.
	•	range iterates over the slice to sum its elements.
	•	Goroutines run the sum calculations in parallel.
	•	Blocking receive from the channel ensures synchronization.
*/

func sum(numList []int, ch chan int) {
	sum := 0
	for index, value := range numList {
		sum += value
		fmt.Println("index:", index, "value:", value, "numList:", numList)
	}
	ch <- sum
}

// NOTE: Regular Channel (UnBuffered)
func channelExample() {
	fmt.Printf("\n# channelExample\r\n")

	numList := []int{1, 2, 3, 4, 5}
	resultChannel := make(chan int)

	go sum(numList[:len(numList)/2], resultChannel)
	go sum(numList[len(numList)/2:], resultChannel)

	firstHalfSum, secondHalfSum := <-resultChannel, <-resultChannel
	close(resultChannel)

	fmt.Println("firstHalf:", numList[:len(numList)/2])
	fmt.Println("secondHalf:", numList[len(numList)/2:])
	fmt.Println("Total Sum:", firstHalfSum+secondHalfSum)
}

// NOTE: Buffered Channel
/*
	Channels can be buffered, meaning they have a fixed size and don’t block until the buffer is full.
	When sending to a buffered channel, the send operation won’t block until the buffer is filled.

	A buffered channel allows sending without waiting for a corresponding receive if the buffer is not full.
*/
func bufferChannelExample() {
	fmt.Printf("\n# bufferChannelExample\r\n")

	// NOTE: Not using buffered channel will cause a deadlock! as the main go routine will be blocked
	newChannel := make(chan int, 2)
	// newChannel := make(chan int) > // ! fatal error: all goroutines are asleep - deadlock!
	newChannel <- 1
	newChannel <- 2
	close(newChannel)
	fmt.Println("firstValue:", <-newChannel)
	fmt.Println("secondValue:", <-newChannel)
}

func unBufferedChannelSecondExample() {
	fmt.Printf("\n# unBufferedChannelSecondExample\r\n")
	ch := make(chan int) // UnBuffered Channel

	// Sender Goroutine
	go func() {
		fmt.Println("Sending value to channel")
		ch <- 42
		fmt.Println("Sending second value to channel (Comes before main logs)")
		ch <- 43
		fmt.Println("This won't be printed")
		close(ch)
	}()

	// Main Goroutine receiving the value
	fmt.Println("Received value:", <-ch)
	fmt.Println("Received value:", <-ch)
}

// NOTE: UnBuffered Channels
/*
	•	UnBuffered channels do not have any capacity to hold values. When a value is sent to an UnBuffered channel,
		the sending Goroutine is blocked until another Goroutine receives the value.
	•	Likewise, a receiving Goroutine is blocked until a value is sent to the channel.
	•	This provides a way to synchronize between Goroutines.
*/

// NOTE: Buffered Channels
/*
	•	Buffered channels have a fixed capacity to store values. Sending to a buffered channel does not block until the buffer is full.
	•	Similarly, receiving from a buffered channel only blocks if there are no values in the buffer.
	•	Buffered channels allow Goroutines to communicate asynchronously up to the buffer size.
*/

// NOTE: Summary
/*
	•	Un‌Buffered channels block both the sender and receiver until they are ready to communicate, making them good for synchronization.
	•	Buffered channels allow values to be sent asynchronously, up to the buffer capacity, which can improve performance by decoupling
		the sending and receiving Goroutines.
*/

// NOTE: Getting the Size of a Channel
/*
	•	len(ch): Returns the current number of elements in the channel.
	•	cap(ch): Returns the total capacity of the buffered channel (the maximum number of elements it can hold).

	Important Notes:

	1.	Unbuffered channels (created without specifying a capacity) will always have a size of 0 when checked with len(), since they
		cannot store elements—data is sent and received directly.
	2.	Using len() on a channel is safe, but be aware that the size may change immediately after you check it due to concurrent operations.
		This means you should not rely on len() for critical logic like synchronization—use synchronization primitives such as sync.Mutex
		or sync.WaitGroup instead if needed.
*/

func main() {
	channelExample()
	bufferChannelExample()
	unBufferedChannelSecondExample()
}
