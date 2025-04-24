package main

import (
	"fmt"
	"time"
)

// Day 33: Select
// the select statement allows a goroutine to wait on multiple communication operations.

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Hello from c1"

	}()

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "Hello from c1 again"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Hello from c2"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1", msg1)
			fmt.Println("Hi")
		case msg2 := <-c2:
			fmt.Println("msg2", msg2)
		}
	}
}
