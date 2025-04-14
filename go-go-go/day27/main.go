package main

import (
	"fmt"
	"reflect"
)

// Day 27: Types and type assertions
// 4 categories of Types: basic, aggregate, reference and interface types.
// Basic types: Numbers, strings, booleans
// Aggregate types: Arrays and structs
// Reference types: Pointers, slices, maps, functions, channels
// Interface types: Define behavior without implementation

func printTyoeInfo(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func main() {
	// 1. Basic types examples
	// remember we dont need to specifiy the type since go will infer it
	var integer = 42
	var floating = 3.14
	var boolean = true
	var str = "Hello, Gopher!"

	fmt.Println("--- Basic types ---")
	fmt.Printf("int: %d (%T)\n", integer, integer)
	fmt.Printf("float64: %f (%T)\n", floating, floating)
	fmt.Printf("bool: %t (%T)\n", boolean, boolean)
	fmt.Printf("string: %s (%T)\n\n", str, str)

	// 2. Aggregate types examples
	// Array (fixed size)
	var arr [3]int = [3]int{1, 2, 3}
	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Alice", 30}

	fmt.Println("--- Aggregate types: ---")
	fmt.Printf("array: %v (%T)\n", arr, arr)
	fmt.Printf("struct: %+v (%T)\n\n", p, p)

	// 3. Reference types examples
	// Slice (dynamic array)
	slice := []int{1, 2, 3}
	// Map
	m := map[string]int{"one": 1, "two": 2}
	// Pointer
	ptr := &integer

	fmt.Println("--- Reference types: ---")
	fmt.Printf("slice: %v (%T)\n", slice, slice)
	fmt.Printf("map: %v (%T)\n", m, m)
	fmt.Printf("pointer: %v (%T) points to %v\n\n", ptr, ptr, *ptr)

	// 4. Interface types and type assertions
	// change this to number or bool to see the difference
	var i interface{} = "Hello, Interface!"

	// Type assertion
	if s, ok := i.(string); ok {
		printTyoeInfo(s)
	} else {
		fmt.Println("Type assertion failed!")
	}

	// Type switch
	fmt.Println("\nType switch example:")
	switch val := i.(type) {
	case int:
		printTyoeInfo(val)
	case string:
		printTyoeInfo(val)
	default:
		fmt.Printf("Unknown type: %v\n", val)
	}

	// Using reflect package for deeper type inspection
	fmt.Println("\nUsing reflect package:")
	fmt.Println("Type of i:", reflect.TypeOf(i))
	fmt.Println("Kind of i:", reflect.TypeOf(i).Kind())
}
