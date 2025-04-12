package main

import (
	"fmt"
	"reflect"
)

// Day 26: Structs
// Structs are user-defined types that helps to create a collection of data
// https://gobyexample.com/structs

type Person struct {
	Name      string
	Age       int
	isWorking bool
}

// newPerson creates a new Person and returns a pointer to it
func newPerson(name string, age int) *Person {
	p := Person{
		Name:      name,
		Age:       age,
		isWorking: false,
	}

	// & means the pointer (memory address) of the p
	return &p
}

func setWorkingStatus(p *Person, status bool) {
	p.isWorking = status
}

func (p *Person) setWorking(status bool) {
	p.isWorking = status
}

func main() {
	// Create a Person using newPerson (returns a pointer)
	gopher := newPerson("Gopher", 15)

	// 3 ways to access and modify a field of a struct:
	// 1. Directly using the pointer to the struct
	// 2. Using a function that takes a pointer to the struct
	// 3. Using a method attached to the struct
	// gopher.isWorking = true
	// setWorkingStatus(gopher, true)
	gopher.setWorking(true)

	// Print the struct using default format
	fmt.Println("Gopher:", gopher)

	// Print the individual fields using format specifiers
	fmt.Printf("Name: %s, Age: %d, isWorking: %t\n", gopher.Name, gopher.Age, gopher.isWorking)

	// %p prints the memory address of the pointer
	fmt.Printf("Memory address of gopher: %p\n", gopher)

	// Create a Person without using a pointer
	penguin := Person{
		Name:      "Penguin",
		Age:       10,
		isWorking: false,
	}

	fmt.Println("\nPenguin:", penguin)
	fmt.Printf("Name: %s, Age: %d, isWorking: %t\n", penguin.Name, penguin.Age, penguin.isWorking)

	// Use &penguin to get the pointer to the struct
	fmt.Printf("Memory address of penguin: %p\n", &penguin)

	// https://pkg.go.dev/reflect#TypeOf
	fmt.Printf("\ngopher - Is pointer? %t\n", reflect.TypeOf(gopher).Kind() == reflect.Ptr)
	fmt.Printf("penguin - Is pointer? %t\n", reflect.TypeOf(penguin).Kind() == reflect.Ptr)
}
