package main

import "fmt"

// Define a struct to represent a "tuple"
type Tuple struct {
	First  int
	Second string
}

func main() {
	// Create instances of the Tuple struct
	tuple1 := Tuple{First: 10, Second: "Hello"}
	tuple2 := Tuple{First: 20, Second: "World"}

	// Access the values using the field names
	fmt.Println("Tuple 1:", tuple1.First, tuple1.Second)
	fmt.Println("Tuple 2:", tuple2.First, tuple2.Second)
}
