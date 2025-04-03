package main

import (
	"errors"
	"fmt"
)

// Day 17: Errors, Panic, and Recover
// https://gobyexample.com/errors
// https://go.dev/blog/error-handling-and-go
// https://go.dev/blog/defer-panic-and-recover
// in Go, nil is the zero value for pointers, interfaces, maps, slices, channles, and functions. it represents the absence of a value or a null reference.
// https://go101.org/article/nil.html

// Errors are returned from functions
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// defer, panic, recover
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Cannot divide by zero")
	}

	fmt.Println("Result:", a/b)
}

func main() {
	fmt.Println("--- Handling Erros ---")

	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	res1, err1 := divide(10, 0)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Result:", res1)
	}

	fmt.Println("\n--- Panic and Recover ---")
	safeDivide(10, 2)
	safeDivide(10, 0)
}
