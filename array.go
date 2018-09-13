package main

import (
	"fmt"
)

func main() {
	var a1 [5]int
	fmt.Println(a1)

	a1[0] = 505
	fmt.Println(a1)

	const size = len(a1)
	fmt.Println(size)

	for i := 0; i < size; i++ {
		fmt.Println(a1[i])
	}

}
