package main

import "fmt"

// Day 15: If/Else statements
// https://gobyexample.com/if-else

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Printf("7 module 2 is: %d.\n", 7%2)

	// change 8 to 9 and see
	if 8%4 == 0 {
		fmt.Println("\n8 is divisible by 4")
	}
	fmt.Printf("8 module 4 is: %d.\n", 8%4)

	// change the num value to see what it prints to console.
	if num := 12; num < 0 {
		fmt.Printf("\n%d is negative", num)
	} else if num < 10 {
		fmt.Printf("\n%d has 1 digit", num)
	} else {
		fmt.Printf("\n%d has multiple digits", num)
	}

	fmt.Println()
}
