package main

// Day 08: Printf general formatting verbs
// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
// https://gobyexample.com/string-formatting

import "fmt"

func main() {
	// focus on general formatting verbs that work with all data types
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign

	var x = 20.5 // default to float64, unless specified declared it as float32
	var txt = "What is general printf verbs?"

	fmt.Printf("Value: %v\t", x)
	fmt.Printf("Value: %#v\t", x)
	fmt.Printf("Value: %v%%\n", x)
	fmt.Printf("Type: %T\n\n", x)
	fmt.Printf("Text: %v\n", txt)
	fmt.Printf("Text: %#v\n", txt)
	fmt.Printf("Text: %v%%\n", txt)
	fmt.Printf("Type: %T\n", txt)

	var num = 100
	fmt.Printf("\nNum - Type: %T, Value: %v, Go-syntax: %#v\n", num, num, num)

	var byte1 byte = 'a'
	var rune1 rune = 'b'
	fmt.Printf("\nbyte1 - Type: %T, Value: %v, Go-syntax: %#v\n", byte1, byte1, byte1)
	fmt.Printf("rune1 - Type: %T, Value: %v, Go-syntax: %#v\n", rune1, rune1, rune1)

	var boolean1 bool = false
	boolean2 := true
	fmt.Printf("\nboolean1 - Type: %T, Value: %v, Go-syntax: %#v\n", boolean1, boolean1, boolean1)
	fmt.Printf("boolean2 - Type: %T, Value: %v, Go-syntax: %#v\n", boolean2, boolean2, boolean2)
}
