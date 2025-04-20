package main

// Day 32: Buffer
// A bytes.Buffer is a temporary storage area in memory to build, modify, or hold byte data efficiently — especially when dealing with I/O, strings, or binary data.
// https://pkg.go.dev/bytes#example-Buffer
// https://www.educba.com/golang-buffer/
// https://www.codingexplorations.com/blog/how-to-use-buffer-in-go-a-comprehensive-guide

import (
	"bytes"
	"fmt"
)

func main() {
	original := "Hello Gopher!"
	fmt.Printf("Original string: %s (type: %T)\n", original, original)

	var b bytes.Buffer
	b.WriteString(original)

	data := make([]byte, 10)
	n, err := b.Read(data)
	if err != nil {
		panic(err)
	}

	readBytes := data[:n]

	fmt.Printf("Read %d bytes: %s (type: %T)\n", n, data, data)
	fmt.Printf("Read bytes: %v (type: %T)\n", readBytes, readBytes)
	fmt.Printf("Converted back to string: %s (type: %T)\n", string(readBytes), string(readBytes))
}
