package main

import "fmt"

// Day 10: printf string and boolean formatting verbs
// strings
// %s prints the value as a string
// %q prints the value as a double-quoted string
// %8s pad with right spaces to 8 characters
// %-8s pad with left spaces to 8 characters
// %x prints the value as hex dump of byte values
// % x prints the value as hex dump with spaces
// boolean
// %t prints the value as a boolean (same as %v)

func main() {
	s := "hello"
	b := true

	fmt.Println("--- String ---")
	fmt.Printf("string - %s\n", s)
	fmt.Printf("double-quoted string - %q\n", s)
	fmt.Printf("right-padded string - %8s\n", s)
	fmt.Printf("left-padded string - %-8s\n", s)
	fmt.Printf("hex dump - %x\n", s)
	fmt.Printf("hex dump with spaces - % x\n", s)

	fmt.Println("\n--- Boolean ---")
	fmt.Printf("boolean with t - %t\n", b)
	fmt.Printf("boolean with v - %v\n", b)
}
