package main

import "fmt"

// Day 21: Arrays
// An array is a collection of elements of the same type with a fixed size that is determined at the time of creation.

func main() {
	// Declare an array of 4 integers (default values = 0)
	var numbers [4]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	// numbers[4] = 50 // This will cause an error: index out of range

	fmt.Println("Array:", numbers)
	fmt.Printf("Type of array: %T\n", numbers)

	// Declare and initialize an array of strings
	languages := [3]string{"Golang", "Python", "TypeScript"}
	fmt.Println("\nLanguages:", languages)
	fmt.Printf("Type of languages array: %T\n", languages)
	// Length of array
	fmt.Println("Length of languages array:", len(languages))

	// Accessing elements
	fmt.Println("\nUsing classic for loop:")
	for index := 0; index < len(languages); index++ {
		fmt.Printf("Language at index %d: %s\n", index, languages[index])
	}

	fmt.Println("\nUsing for...range loop:")
	for index, name := range languages {
		fmt.Printf("Language at index %d: %s\n", index, name)
	}
}
