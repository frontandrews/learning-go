package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{X: 5, Y: 10}
	p2 := Point{X: 5, Y: 10}
	p3 := Point{X: 3, Y: 8}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false
}
