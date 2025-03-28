package main

// Day 11: Printf float formatting verbs
// %e - scientific notation with 'e' as exponent
// %f - decimal point, no exponent
// %.2f - default width, precision 2
// %6.2 - width 6, precision 2
// %g - exponent as needed, only necessary digits

import "fmt"

func main() {
	var pi = 3.14159265359

	fmt.Println("--- Printf float formatting verbs ---")
	fmt.Printf("Scientific notion - %e\n", pi)
	fmt.Printf("Decimal point - %f\n", pi)
	fmt.Printf("With precision 2 - %.2f\n", pi)
	fmt.Printf("With precision 5 - %.5f\n", pi)
	fmt.Printf("With precision 10 - %.10f\n", pi)
	fmt.Printf("Width 6 - %6.2f\n", pi)
	fmt.Printf("Width 10 - %10.2f\n", pi)
	fmt.Printf("Width 20 - %20.2f\n", pi)
	fmt.Printf("Necessary digits - %g\n", pi) // try to change pi variable to longer number and see how this displays
}
