package main

import (
	"the-gopher-learning/go-go-go/day19/greetings"
)

// Day 19: Packages, imports, and exports
// How to create a package in go: https://www.golang-book.com/books/intro/11
// go library: https://pkg.go.dev/std

func main() {
	greetings.Hello("Gopher")
	greetings.Goodbye("Gopher")

	// try to type greetings. this should give you a list of all the functions in the package
	// tryMe is not on the list
}
