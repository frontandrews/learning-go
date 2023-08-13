package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	fmt.Println("Full Name:", person.GetFullName())
}
