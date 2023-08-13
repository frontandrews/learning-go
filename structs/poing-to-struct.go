package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{X: 5, Y: 10}
	fmt.Println("Original Point:", p)

	// Using a pointer to modify the struct's values
	pointerToP := &p
	pointerToP.X = 15
	fmt.Println("Modified Point:", p)
}
