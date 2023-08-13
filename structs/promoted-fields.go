package main

import "fmt"

type Address struct {
    Street string
    City   string
}

type Person struct {
    FirstName string
    LastName  string
    Age       int
    Address
}

func main() {
    p := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        Address: Address{
            Street: "123 Main St",
            City:   "Anytown",
        },
    }

    fmt.Println("Full Name:", p.FirstName, p.LastName)
    fmt.Println("City:", p.City)
}
