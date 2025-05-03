package main

import (
	"fmt"
	"runtime"
	"time"
)

// Day 35: Scheduler
// Go scheduler allows us to nuderstand more deeply how goroutines are scheduled.

func main() {
	runtime.GOMAXPROCS(1) // use 1 CPU core

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1:", i)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2 ==>", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second) // wait for goroutines to finish
}
