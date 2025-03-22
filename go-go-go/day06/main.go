package main

// Day 06: Data Type - string

import "fmt"

func main() {
	var s1 string = "Hello, World!"
	var s2 = "Go Go Go!"
	var s3 string
	s4 := "Hello, Gopher!"
	const s5 string = "Hello, Gopher! (this is declared with const)"

	fmt.Println("String----------")
	fmt.Println("string 1", s1)                                      // Returns Hello, World!
	fmt.Println("string 2", s2)                                      // Returns Hello, Go!
	fmt.Println("string 3" + s3 + " (this returns an empty string)") // Returns empty string
	fmt.Println("string 4", s4)                                      // Returns Hello, Gopher!
	fmt.Println("string 5", s5)                                      // Returns Hello, Gopher! (this is declared with const)
}
