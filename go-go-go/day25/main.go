package main

import (
	"fmt"
)

// Day 25: Make()
// a built-in function make, help to create and initialize slices, maps and channels
// make([]Type, length, capacity)

func main() {
	// Create a slice of int with length 3 and capacity 5
	// make([]type, lenght, capacity)
	// length = the number of elemets in the slice
	// capacity = the number of elements the slice can grow to
	numbers := make([]int, 3, 5)
	fmt.Println("Slice:", numbers)
	fmt.Println("Capacity: ", cap(numbers))
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	// Capacity is not a hard limit, it ca grow if needed
	fmt.Println("Slice after added 6th element:", numbers)
	fmt.Println("Capacity: ", cap(numbers))

	// Create a map with string keys and int values
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Println("\nMap:", ages)

	// Create a channel of int
	ch := make(chan int, 2) // buffered channel with capacity 2
	ch <- 10
	fmt.Println("\nSent 10")
	ch <- 20
	fmt.Println("Sent 20")

	// ❌ This will cause a deadlock because the channel is full and no one is receiving
	// Uncomment the below and see what happens
	// ch <- 30

	fmt.Println("Receiving:", <-ch)
	fmt.Println("Receiving:", <-ch)

	// ✅ Now the channel is empty, so we can send again
	ch <- 30
	fmt.Println("Sent 30")

	fmt.Println("Receiving:", <-ch)
}
