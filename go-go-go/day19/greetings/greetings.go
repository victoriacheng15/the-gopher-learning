package greetings

import "fmt"

func Hello(name string) {
	fmt.Println("Hello, ", name)
}

func Goodbye(name string) {
	fmt.Println("Goodbye, ", name)
}

// Need to be capitalized to be exported
func tryMe() {
	fmt.Println("Try me!")
}
