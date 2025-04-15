package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Day 29: Context
// context package provides a standard way to solve the problem of managing the state during a request.
// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
// https://gobyexample.com/context
// net/http package: https://pkg.go.dev/net/http

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// change Second to Millisecond and see what happens
	time := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()

	url := "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("✅ Got post:")
	fmt.Printf("ID: %d, Title: %s\n", post.ID, post.Title)
}
