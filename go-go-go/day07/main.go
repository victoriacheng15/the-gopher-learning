package main

import "fmt"

// Day 07: Data Type - byte and rune
// Byte vs Rune
// byte - 8-bit values, often representing ASCII characters.
// https://www.ascii-code.com/
// rune - 32-bit values, representing Unicode code points.
// https://symbl.cc/en/unicode-table/

func main() {
	var b1 byte = 'a'
	var b2 = 'b'
	var b3 byte
	b4 := 'c'
	// this will show as a string, not a byte
	b5 := "d"

	fmt.Println("--- Bytes ---")
	fmt.Println("b1", b1) // Output: 97
	fmt.Println("b2", b2) // Output: 98
	fmt.Println("b3", b3) // Output: 0
	fmt.Println("b4", b4) // Output: 99
	fmt.Println("b4", b5) // Output: d

	var r1 rune = 'a'
	var r2 = 'b'
	var r3 rune
	r4 := 'c'
	// this will show as a string, not a rune
	r5 := "d"

	fmt.Println("\n--- Runes ---")
	fmt.Println("r1", r1) // Output: 97
	fmt.Println("r2", r2) // Output: 98
	fmt.Println("r3", r3) // Output: 0
	fmt.Println("r4", r4) // Output: 99
	fmt.Println("r5", r5) // Output: d

	// Noticed a couple of things:
	// If hover to b1, it will show byte as the type
	// while hover to b2, it will show rune as the type
	// For rune, all variables excluded r5 are of type rune
	// If use double quotes for btye and rune, it will show the type as string
	// Keep in this mind, if want to delcared the value as either rune or byte, use single quotes
}
