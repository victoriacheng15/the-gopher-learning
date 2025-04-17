package main

import (
	"fmt"
	"time"
)

// Day 31: Channels
// channels are the pipes that connect concurrent goroutines.
// The channel operator <- is used to send or receive values through the channel.
// https://golangbot.com/channels/

func main() {
	ch := make(chan string, 2)

	ch <- "1st message"
	ch <- "2nd message"

	// At this point, the channel buffer is full.
	// Sending a 3rd message now would block the program:
	// ch <- "3rd message" // ❌ blocks forever if no receiver

	// We use a goroutine to send the 3rd message later.
	// The 'go' keyword runs the function in the background (asynchronous),
	// so it does NOT wait for the message to be received right away.
	// It allows the main function to keep running while this waits its turn.
	go func() {
		// try to change 1 to 2 and see what happens
		time.Sleep(1 * time.Second)
		ch <- "3rd message"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "4th message"
	}()

	// Receiving messages from the channel.
	// This will free up space in the buffer, allowing the goroutine to send.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // after 1st delay
	fmt.Println(<-ch) // after 1st delay

	// Example Analogy:
	// A channel is a waiting room with 2 chairs.
	// You have 4 people (messages) coming in.
	// The first two sit right away.
	// The next two (sent by goroutines) will stand in line until a chair becomes free.
}
