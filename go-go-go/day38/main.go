package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Person represents a basic person structure for JSON examples
type Person struct {
	Name string `json:"name"` // Field appears in JSON as "name"
	Age  int    `json:"age"`  // Field appears in JSON as "age"
}

// PersonWithExtra extends Person to handle unknown fields
type PersonWithExtra struct {
	Person
	Extra map[string]interface{} `json:"-"` // The "-" tag prevents this field from being used in JSON
}

func goToJson() {
	p := Person{Name: "Alice", Age: 30}

	// Convert Go struct to pretty-printed JSON
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Println(string(jsonData))
}

func jsonToGo() {
	jsonData := `{"name": "Bob", "age": 25}`

	// Convert JSON to Go struct
	var p Person
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func handleUnknownFields() {
	jsonData := `{"name": "Charlie", "age": 28, "city": "New York"}`

	// First unmarshal to capture all data
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &raw); err != nil {
		log.Fatalf("Initial JSON unmarshaling failed: %s", err)
	}

	// Create our person with extra fields
	var p PersonWithExtra
	p.Extra = make(map[string]interface{})

	// Process known and unknown fields
	for k, v := range raw {
		switch k {
		case "name":
			p.Name = v.(string)
		case "age":
			p.Age = int(v.(float64)) // JSON numbers are float64
		default:
			p.Extra[k] = v
		}
	}

	fmt.Printf("Person: %+v\n", p.Person)
	fmt.Printf("Extra fields: %v\n", p.Extra)
}

func strictDecoding() {
	jsonData := `{"name": "Dave", "age": 35, "country": "Canada"}`

	var p Person
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&p); err != nil {
		log.Printf("Strict decoding failed: %v", err)
	} else {
		fmt.Printf("Strict decoded: %+v\n", p)
	}
}

func main() {
	fmt.Println("==== Go -> JSON ====")
	goToJson()

	fmt.Println("\n==== JSON -> GO ====")
	jsonToGo()

	fmt.Println("\n==== Handle Unknown Fields ====")
	handleUnknownFields()

	fmt.Println("\n==== Strict Decoding ====")
	strictDecoding()
}
