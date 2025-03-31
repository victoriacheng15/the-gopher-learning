package main

import "fmt"

// Day 14: Range - iterates over elements in a variety of data structures
// Range iteraates over elements in a variety of dataa structures
// https://gobyexample.com/range

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	fmt.Println("--- Iterating over nums ---")
	for index, num := range nums {
		// if dont use index, can write _, num := range nums
		sum += num
		fmt.Printf("The number on index %d is %d and sum is %d\n", index, num, sum)
	}
	fmt.Println("Final sum:", sum)

	fmt.Println("\n--- Iterating over key/value pairs in a map ---")
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println("\n--- Iterating over the keys in a map ---")
	for key := range kvs {
		fmt.Println("key:", key)
	}

	fmt.Println("\n--- Iterating over the values in a map ---")
	for _, value := range kvs {
		fmt.Println("Value:", value)
	}

	fmt.Println("\n--- Iterating over a string ---")
	fmt.Println("index rune string")
	for i, c := range "go" {
		fmt.Println(i, c, string(c))
		// https://symbl.cc/en/0067/ -> for letter g
		// https://symbl.cc/en/004F/ -> for letter o
	}
}
