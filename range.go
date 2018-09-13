package main

import (
	"fmt"
)

func main() {

	var myArray = [4]int{1, 2, 3, 4}

	sum := 0

	for index, _ := range myArray {
		sum += myArray[index]
	}

	fmt.Println(sum)

}
