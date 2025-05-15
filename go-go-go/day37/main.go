package main

import "fmt"

// Day 37: Pointers
// a feature in Go that allows you to work with memory addresses directly.
// https://gobyexample.com/pointers

func zeroValue(ival int) {
	fmt.Println("Inside zeroValue, before change:", ival)
	ival = 0
}

func zeroPointer(iptr *int) {
	if iptr == nil {
		fmt.Println("iptr is nil")
		return
	}
	*iptr = 0
}

func main() {
	num := 5
	fmt.Println("Initial: ", num)

	zeroValue(num)
	fmt.Println("zeroValue:", num)

	zeroPointer(&num)
	fmt.Println("zeroPointer:", num)

	fmt.Println("Pointer:", &num)
}
