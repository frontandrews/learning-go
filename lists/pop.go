package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	numbers = numbers[:len(numbers)-1]

	fmt.Println(numbers) // [1 2 3]
}
