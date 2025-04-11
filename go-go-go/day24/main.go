package main

import "fmt"

// Day 24: Maps
// Maps are a data structure that allows you to store key-value pairs.
// https://gobyexample.com/maps

func setupPerson() map[string]string {
	// Declare and initialize a map
	person := map[string]string{}

	fmt.Printf("Type of map: %T\n", person)
	fmt.Println("Empty map:", person)

	person["name"] = "Gopher"
	person["age"] = "15"
	person["city"] = "Toronto"
	person["country"] = "Canada"
	person["occupation"] = "Software Engineer"
	person["hobby"] = "Photography"

	return person
}

func iteratingMap(person map[string]string) string {
	result := ""

	for key, value := range person {
		result = fmt.Sprintf("Key: %s, Value: %s\n", key, value)
	}

	return result
}

type KeyValue struct {
	Key    string
	Value  string
	Exists bool
}

func iteratingSpecifcOrder(person map[string]string) []KeyValue {
	specificOrder := []string{"name", "age", "city", "country", "occupation", "hobby"}
	var result []KeyValue

	for _, key := range specificOrder {
		value, exists := person[key]
		result = append(result, KeyValue{Key: key, Value: value, Exists: exists})
	}

	fmt.Println("Result:", result)
	return result
}

func main() {
	person := setupPerson()

	fmt.Println("\nMap with values:", person)
	fmt.Println("Length of map:", len(person))

	// Accessing values
	message := fmt.Sprintf("The name is %s and %s years old", person["name"], person["age"])
	fmt.Println("\nMessage: ", message)

	// Delete a key
	delete(person, "hobby")
	fmt.Println("\nAfter deleting hobby:", person)
	fmt.Println("Length of map:", len(person))

	// Check if a key exists
	fmt.Println("\nChecking if age key exists:")
	value, exists := person["age"]
	if exists {
		fmt.Println("Age:", value)
	} else {
		fmt.Println("Age not found")
	}

	// check if a key exists
	fmt.Println("\nChecking if hobby key exist:")
	value, exists = person["hobby"]
	if exists {
		fmt.Println("Hobby:", value)
	} else {
		fmt.Println("Hobby not found")
	}

	// Iterating over a map
	fmt.Println("\nIterating over map:")
	result := iteratingMap(person)
	fmt.Println(result)

	// Iterating over a map with a specific order
	fmt.Println("Iterating over map with a specific order:")
	orderedResults := iteratingSpecifcOrder(person)
	for _, item := range orderedResults {
		if item.Exists {
			fmt.Printf("Key: %s, Value: %s\n", item.Key, item.Value)
		} else {
			fmt.Printf("Key: %s does not exist\n", item.Key)
		}
	}

	// Clearing a map
	clear(person)
	fmt.Println("\nAfter clearing the map:", person)
}
