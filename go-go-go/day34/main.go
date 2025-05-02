package main

import (
	"fmt"
	"sync"
)

// Day 34: Mutex
// In Go, we can run code at the same time using goroutines.
// But when multiple goroutines access the same data, it can cause *race conditions* (unexpected behavior).
//
// A *mutex* (mutual exclusion) is a tool from the `sync` package.
// It helps us protect shared data by allowing only one goroutine to access it at a time.
//
// https://www.sohamkamani.com/golang/mutex/
// https://pkg.go.dev/sync

// Run this program with and without the -race flag to see the difference:
// go run -race main.go
//
// The -race flag enables Go's race detector, which checks for unsafe access
// to shared variables between goroutines. It slows down the program a bit,
// so race conditions may not always change the final result.
//
// Try running without -race and you'll often see the final counter is less than expected,
// showing that the goroutines are interfering with each other.
//
// With -race, the final count may still look correct (like 1000),
// but Go will warn you about unsafe concurrent access (data race).

var counter int
var maxRange int = 1000

// WithoutMutex shows race condition
func WithoutMutex() {
	counter = 0
	var wg sync.WaitGroup

	for i := 0; i < maxRange; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Without mutex - Final Counter:", counter)
}

// WithMutex prevents race condition
func WithMutex() {
	counter = 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < maxRange; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("With mutex - Final Counter:", counter)
}

func main() {
	fmt.Println("Running without mutex:")
	WithoutMutex()

	fmt.Println("\nRunning with mutex:")
	WithMutex()
}
