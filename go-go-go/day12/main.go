package main

import "fmt"

// Day 12: Printf vs Sprintf
// https://pkg.go.dev/fmt
// https://faun.pub/golangs-fmt-sprintf-and-printf-demystified-4adf6f9722a2
// https://dev.to/robogeek95/exploring-the-go-fmt-package-1p44
// Sprintf - write formatted output to a string, use when you need to store formatted date in a string, such as for custom error messages
// Printf - print formatted output, use when you want to display formatted text to the console, such as for debugg

func main() {
	fmt.Println("--- Sprintf ---")
	name := "Penguin"
	age := 34
	message := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	// %s verb for strings, %d for integers
	fmt.Println(message)

	fmt.Println("\n--- Printf ---")
	product := "Linux Laptop"
	price := 1500.59
	fmt.Printf("The %s costs $%.2f.\n", product, price)
	// %s verb for strings, %.2f for float with 2 decimal places

	// Hover over Sprintf and Printf to see their full function signatures.
	// Sprintf format a string and return a string (the formatted result)
	// Printf format a string and write the result to standard output,
	// returning the number of bytes written (int) and an (error)
}
