package main

import "fmt"

// Day 18: Functions
// https://gobyexample.com/functions

func plus(a int, b int) int {
	// include types for arguments and return value
	return a + b
}

func square(a int) int {
	return a * a
}

func sumAndDiff(a int, b int) (int, int) {
	// multiple return values
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	var num1 = 1
	num2 := 2
	num3 := 6

	plusRes := plus(num1, num2)
	plusResString := fmt.Sprintf("%d + %d = %d", num1, num2, plusRes)
	fmt.Println(plusResString)

	squareRes := square(num3)
	squareResString := fmt.Sprintf("%d * %d = %d", num3, num3, squareRes)
	fmt.Println(squareResString)

	sum, diff := sumAndDiff(num2, num3)
	fmt.Printf("Sum and Difference of %d and %d: Sum = %d, Difference = %d\n", num2, num3, sum, diff)
}
