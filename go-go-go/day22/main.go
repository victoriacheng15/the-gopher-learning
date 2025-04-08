package main

import "fmt"

// Day 22: Arrays
// An array is a collection of elements of the same type with a fixed size that is determined at the time of creation.

func main() {
	// Declare an array of 4 integers (default values = 0)
	var numbers [4]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	// numbers[4] = 50
	// Attempting to assign a value to numbers[4] will cause an error: index out of range.
	// This is because the array 'numbers' has a fixed size of 4, and valid indices are 0, 1, 2, and 3.

	fmt.Println("Array:", numbers)
	fmt.Printf("Type of array: %T\n", numbers)

	// The size of the array is inferred from the number of elements provided (3 elements in this case).
	languages := [3]string{"Golang", "Python", "TypeScript"}
	fmt.Println("\nLanguages:", languages)
	fmt.Printf("Type of languages array: %T\n", languages)
	// Length of array
	fmt.Println("Length of languages array:", len(languages))

	// Iterating over the array using indices
	fmt.Println("\nUsing classic for loop:")
	for index := 0; index < len(languages); index++ {
		fmt.Printf("Language at index %d: %s\n", index, languages[index])
	}

	// The for...range loop is a more idiomatic way to iterate over arrays in Go, providing both the index and value.
	fmt.Println("\nUsing for...range loop:")
	for index, name := range languages {
		fmt.Printf("Language at index %d: %s\n", index, name)
	}
}
