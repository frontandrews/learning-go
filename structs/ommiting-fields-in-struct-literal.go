package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{X: 5}
	fmt.Println("Point:", p) // Point: {5 0}
}
