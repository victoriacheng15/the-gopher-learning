package main

import "fmt"

// Day 36: Generics
// Generic is a feature in Go that allows you to write functions and data structures that can operate on different types without sacrificing type safety.

// Generic struct: holds any type of value
type Box[T any] struct {
	Value T
}

func FirstElement[T any](items []T) T {
	return items[0]
}

func main() {
	// Use the generic function
	fmt.Println("Int:", FirstElement([]int{10, 20, 30}))
	fmt.Println("String:", FirstElement([]string{"a", "b", "c"}))
	fmt.Println("Float64:", FirstElement([]float64{1.1, 2.2, 3.3}))
	fmt.Println("Bytes:", FirstElement([]byte{'a', 'b', 'c'}))

	// Use the generic struct
	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "hello, Gopher!"}
	byteBox := Box[byte]{Value: 'z'}

	fmt.Println("\nInt box:", intBox.Value)
	fmt.Println("String box:", stringBox.Value)
	fmt.Println("Byte box:", byteBox.Value)
}
