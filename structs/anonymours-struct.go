package main

import "fmt"

func main() {
	person := struct {
		FirstName string
		LastName  string
		Age       int
	}{
		FirstName: "Alice",
		LastName:  "Johnson",
		Age:       28,
	}

	fmt.Println("Full Name:", person.FirstName, person.LastName)
}
