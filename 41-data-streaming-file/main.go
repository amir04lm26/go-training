package main

import (
	"bufio"
	"fmt"
	"os"
)

// NOTE: Data Streaming
/*
	A common practice for handling large volumes of data efficiently. Streaming is used when you want to process data
	incrementally instead of loading everything into memory at once. It’s especially useful when working with files,
	network communication, or large datasets.
*/

// NOTE: Data Streaming in Go
/*
	Go provides several ways to work with streaming data, such as reading and writing from files, network connections,
	and HTTP responses.
*/

// NOTE: Streaming File Data
/*
	Reading a file in chunks allows you to process large files without loading the entire file into memory.
*/

// NOTE: bufio
/*
	Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object
	(Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
*/

// NOTE: When to Use bufio?
/*
	•	Large files: Reading or writing large files where byte-by-byte operations would be slow.
	•	Network communication: Reduces the number of I/O system calls, making it more efficient to send or receive data over the network.
	•	Line-by-line processing: Ideal when you need to read or process input one line at a time, such as logs.

	bufio helps by buffering data, minimizing I/O calls, and providing convenient methods for handling common tasks like reading lines or
	writing strings.
*/

// NOTE: Explanation
/*
	•	os.Open: Opens the file for reading.
	•	bufio.NewReader: Creates a buffered reader to read the file in chunks.
	•	buffer[:n]: Processes each chunk of data as it’s read from the file.
	•	defer file.Close(): Ensures that the file is closed after reading is complete.

	This approach ensures that even very large files can be processed without consuming too much memory.
*/

func main() {
	file, err := os.Open("gtp-resources.txt") // NOTE: Related to the script execution dir
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println(reader.Size())
	buffer := make([]byte, 1024) // NOTE: Read 1KB at a time
	for {
		numberOfBytes, err := reader.Read(buffer)
		if numberOfBytes > 0 {
			fmt.Println(string(buffer[:numberOfBytes]))
		}
		if err != nil {
			break // End of file
		}
	}
}
