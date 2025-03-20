package main

import (
	"fmt"
	"math"
)

// Day 03: Declaring variables with types using `const`
// 'const' is similar to 'var' for declaring variables.
// The difference is that variables declared with 'var' can be re-assigned,
// while variables declared with 'const' cannot be re-assigned.

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const test string = "try to re-assign me"
	// This line will cause a compilation error
	// Cannot re-assign a value to a const
	// uncomment the below and run 'go run main.go'
	// test = "this will fail"
	fmt.Println(test)
}
