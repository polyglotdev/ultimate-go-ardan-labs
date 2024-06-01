package main

import (
	"fmt"
)

// reader is an interface that defines the act of reading data.
type reader interface {
	// read reads up to len(b) bytes into b. It returns the number of bytes read (0 <= n <= len(b)) and any error encountered.
	// Even if read returns n < len(b), it may use all of b as scratch space during the call.
	// If some data is available but not len(b) bytes, read conventionally returns what is available instead of waiting for more.
	// When read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read.
	// It may return a non-nil error or nil error if n > 0. An error is returned only if no bytes were read.
	read(b []byte) (int, error)
}

// file defines a system file.
type file struct {
	name string
}

// read implements the reader interface for a file.
// A function in Go programming language to read data into a byte array from a predefined string.
func (f file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	n := copy(b, s)
	return n, nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (p pipe) read(b []byte) (int, error) {
	s := `{name: "elijah", title: "software engineer"}`
	n := copy(b, s)
	return n, nil
}

func main() {
	f := file{"data.xml"}
	data := make([]byte, 100)
	n, err := f.read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Buffer Content: ", string(data[:n]))
	fmt.Println("Length of data: ", len(data))
	fmt.Println("Bytes Read: ", n)

	// For loop to print each byte in the buffer along with its index, size, length and value
	for i, b := range data[:n] {
		fmt.Printf("Index: %d, Size: %d bytes, Length: %d, Value: %c\n", i, 1, len(data[:n]), b)
	}

	p := pipe{"data.pipe"}
	data = make([]byte, 20)

	n, err = p.read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	println("----------------------------------------------------------")

	fmt.Println("Buffer Content: ", string(data[:n]))
	fmt.Println("Length of data: ", len(data))
	fmt.Println("Bytes Read: ", n)

	// For loop to print each byte in the buffer along with its index, size, length and value
	for i, b := range data[:n] {
		fmt.Printf("Index: %d, Size: %d bytes, Length: %d, Value: %c\n", i, 1, len(data[:n]), b)
	}

}
