package main

import "fmt"

// Day 23: Slices
// Just like arrays, but more flexible — you can add or remove elements.
// Arrays have a fixed size, so you can’t add new values once they’re full.

func main() {
	// Declare and initialize a slice of integers
	numbers := []int{10, 20, 30, 40}

	// Append more values — append returns a new slice with the added elements
	numbers = append(numbers, 50)

	fmt.Println("Numbers slice:", numbers)
	fmt.Printf("Type of slice: %T\n", numbers)
	fmt.Println("Length of slice:", len(numbers))
	fmt.Println("Capacity of slice:", cap(numbers))

	// Declare and initialize a slice of strings
	languages := []string{"Python", "Golang", "JavaScript"}
	languages = append(languages, "TypeScript")

	fmt.Println("\nLanguages slice:", languages)
	fmt.Printf("Type of languages slice: %T\n", languages)
	fmt.Println("Length of languages slice:", len(languages))
	fmt.Println("Capacity of languages slice:", cap(languages))

	// Accessing elements using classic for loop
	fmt.Println("\nUsing classic for loop:")
	for i := 0; i < len(languages); i++ {
		fmt.Printf("nLanguage at index %d: %s\n", i, languages[i])
	}

	// Accessing elements using range
	fmt.Println("\nUsing for...range loop:")
	for i, name := range languages {
		fmt.Printf("nLanguage at index %d: %s\n", i, name)
	}
}
