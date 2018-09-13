package main

import (
	"fmt"
)

func changeValue(value *int) {
	*value *= 2

	fmt.Println(*value)
}

func main() {

	value := 2
	fmt.Println(value)
	changeValue(&value)
	fmt.Println(value)
}
