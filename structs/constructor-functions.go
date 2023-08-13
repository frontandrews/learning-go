package main

import (
	"fmt"
)

// Define a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Constructor function for creating a new Person instance
func NewPerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	// Create a new Person instance using the constructor function
	person1 := NewPerson("Alice", "Johnson", 28)
	person2 := NewPerson("Bob", "Smith", 35)

	// Access the fields of the Person instances
	fmt.Println("Person 1:", person1.FirstName, person1.LastName, person1.Age)
	fmt.Println("Person 2:", person2.FirstName, person2.LastName, person2.Age)
}
