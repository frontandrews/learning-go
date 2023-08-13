package main

import "fmt"

// Person struct represents a person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Implementing the String() method of the Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Name: %s %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	// Create a person instance
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Printing the person instance using fmt.Println
	fmt.Println(p) // Output: Name: John Doe, Age: 30
}
