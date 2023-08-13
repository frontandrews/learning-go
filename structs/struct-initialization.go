package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{
		X: 5,
		Y: 10,
	}
	fmt.Println("Point:", p)
}
