package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rectangle := &Rectangle{Width: 5, Height: 10}
	fmt.Println("Area:", rectangle.Area())
}
