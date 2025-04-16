package main

import (
	"fmt"
	"time"
)

// Day 30: Goroutines
// goroutines allow us to write concurrent programs in go.

func count(from string) {
	for i := 0; i < 5; i++ {
		// sleep for 1 second
		time.Sleep(1 * time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {
	count("direct")

	go count("Goroutine")

	// goroutine from annoymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("Is it going...?")

	// Each iteration of the loop takes 1 second,
	// and since there are 5 iterations, the total execution time for the loop is 5 seconds.
	// The main function only sleeps for a certain amount of time.
	// By changing the sleep time (between 1 to 5 seconds), you can observe how the loop progresses.
	// If the sleep time is less than 5 seconds, the loop will only complete the number of iterations
	// that fit within the given sleep time, and any remaining iterations will be executed after the sleep time ends.
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
