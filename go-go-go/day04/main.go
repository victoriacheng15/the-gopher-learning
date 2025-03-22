package main

import "fmt"

// Day 04: Data types - Booleans and Integers

func allBooleans() {
	// boolean
	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	fmt.Println("Boolean")
	fmt.Println("boolean 1: ", b1) // Returns true
	fmt.Println("boolean 2: ", b2) // Returns true
	fmt.Println("boolean 3: ", b3) // Returns false
	fmt.Println("boolean 4: ", b4) // Returns true
}

func allSignedIntger() {
	var n1 int = 500
	var n2 int = -4500
	n3 := 100
	var n4 int

	fmt.Println("\nInteger - signed integers")
	fmt.Println("Integer 1: ", n1) // Type: int, Value: 500
	fmt.Println("Integer 2: ", n2) // Type: int, Value: -4500
	fmt.Println("Integer 3: ", n3) // Type: int, Value: 100
	fmt.Println("Integer 4: ", n4) // Type: int, Value: 0

	// Signed Integer Types:
	// int    - platform-dependent:
	//         32 bits on 32-bit systems (-2147483648 to 2147483647)
	//         64 bits on 64-bit systems (-9223372036854775808 to 9223372036854775807)
	// int8   - 8 bits / 1 byte  → Range: -128 to 127
	// int16  - 16 bits / 2 bytes → Range: -32768 to 32767
	// int32  - 32 bits / 4 bytes → Range: -2147483648 to 2147483647
	// int64  - 64 bits / 8 bytes → Range: -9223372036854775808 to 9223372036854775807
}

func allUnsignedInteger() {
	// Unsigned integers
	var n1 uint = 500
	var n2 uint = 4500
	n3 := 100
	var n4 uint

	fmt.Println("\nInteger - unsigned integers")
	fmt.Println("Integer 1: ", n1) // Type: uint, Value: 500
	fmt.Println("Integer 2: ", n2) // Type: uint, Value: 4500
	fmt.Println("Integer 3: ", n3) // Type: uint, Value: 100
	fmt.Println("Integer 4: ", n4) // Type: uint, Value: 0

	// Unsigned Integer Types:
	// uint    - platform-dependent:
	//           32 bits on 32-bit systems (0 to 4294967295)
	//           64 bits on 64-bit systems (0 to 18446744073709551615)
	// uint8   - 8 bits / 1 byte  → Range: 0 to 255
	// uint16  - 16 bits / 2 bytes → Range: 0 to 65535
	// uint32  - 32 bits / 4 bytes → Range: 0 to 4294967295
	// uint64  - 64 bits / 8 bytes → Range: 0 to 18446744073709551615
}

func main() {
	allBooleans()
	allSignedIntger()
	allUnsignedInteger()
}
