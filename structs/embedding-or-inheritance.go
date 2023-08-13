package main

import (
	"fmt"
)

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address // Embedded struct
}

func main() {
	person := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       25,
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			ZipCode: "12345",
		},
	}

	fmt.Println("Full Name:", person.FirstName, person.LastName)
	fmt.Println("Address:", person.Address.Street, person.Address.City, person.Address.ZipCode)
}
