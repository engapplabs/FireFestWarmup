package main

import (
	"fmt"
)

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(sumNumbers(1, 2, 34, 5, 6))
	fmt.Println(sumNumbers(1, 2))
}
