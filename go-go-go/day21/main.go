package main

import "fmt"

// Day 21: Type inference
// type inference gives go the ability to infer the type of a variable based on the value assigned to it.
// which is why it is possible to declare variables without specifying the type.

func main() {
	var name = "Gopher"
	fmt.Printf("Variable 'name' is of type %T\n", name)

	var age = 42
	fmt.Printf("Variable 'age' is of type %T\n", age)

	pi := 3.14159
	fmt.Printf("Variable 'pi' is of type %T\n", pi)
}
