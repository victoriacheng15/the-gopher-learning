package main

import "fmt"

// Day 02: Declaring variables with types using `var`
// https://gobyexample.com/variables

func main() {
	// declare 1 or more variable with var
	var init = "initial"
	fmt.Println(init)

	// declare multiple variables at once
	var num1, num2 int = 1, 2
	fmt.Println(num1, num2)

	// infer type based on initialized variable
	var boolean = true
	fmt.Println(boolean)

	// variables declared without an initialization are zero-valued
	var empty int
	fmt.Println(empty)

	// can initialize with :=
	shorthand := "gopher"
	fmt.Println(shorthand)
}
