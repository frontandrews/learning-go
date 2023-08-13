package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var result []int

	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	fmt.Println(result) // [2 4]
}
