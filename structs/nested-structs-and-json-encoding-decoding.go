package main

import (
	"encoding/json"
	"fmt"
)

// Address struct represents an address
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

// Person struct represents a person
type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}

func main() {
	// Create a nested struct instance
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}

	// Encoding the struct to JSON
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	fmt.Println("Encoded JSON:")
	fmt.Println(string(jsonData))

	// Decoding JSON back to a struct
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded Struct:")
	fmt.Println(decodedPerson)
}
