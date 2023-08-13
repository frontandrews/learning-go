package main

import "fmt"

// Person struct with public (exported) and private fields
type Person struct {
	FirstName string // Public field (exported)
	lastName  string // Private field
}

// GetName returns the full name of the person
func (p Person) GetName() string {
	return p.FirstName + " " + p.lastName
}

func main() {
	p := Person{
		FirstName: "John",
		lastName:  "Doe",
	}

	fmt.Println("Full Name:", p.GetName()) // Using a method to access private fields
}
