package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	p := Person{
		FirstName: "James",
		LastName:  "Smith",
	}

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, field.Tag)
	}
}
