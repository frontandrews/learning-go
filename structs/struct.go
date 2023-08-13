package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Creating an instance of the struct
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Accessing struct fields
	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
	fmt.Println("Age:", person.Age)
}
