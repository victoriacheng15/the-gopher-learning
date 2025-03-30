package main

import "fmt"

// Day 13: For Loop
// https://gobyexample.com/for
// https://yourbasic.org/golang/for-loop/

func main() {

	fmt.Println("With a condition")
	i := 1
	for i <= 5 {
		fmt.Printf("%d\t", i)
		i++
	}

	fmt.Println("\n\nClassic initial / condition / after loop")
	for j := 0; j < 5; j++ {
		fmt.Printf("%d\t", j)
	}

	fmt.Println("\n\nDo this n times with range")
	for i := range 5 {
		fmt.Printf("%d\t", i)
	}

	fmt.Println("\n\nInfinite loop with break statement")
	for {
		fmt.Println("this will run forever")
		// break will stop this infinite loop
		break
	}

	fmt.Println("\nLooping through an  with continue statement")
	for n := range 8 {
		if n == 3 || n == 5 {
			// skipping 3 and 5
			continue
		}
		fmt.Printf("%d\t", n)
	}

	fmt.Println()
}
