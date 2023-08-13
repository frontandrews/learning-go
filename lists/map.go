package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := make([]int, len(numbers))

	for i, num := range numbers {
		result[i] = num * 2
	}

	fmt.Println(result) // [2 4 6 8 10]
}
