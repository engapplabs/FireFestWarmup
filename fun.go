package main

import (
	"fmt"
)

func getDouble(number int) int {
	return number * 2
}

func changeVar(a *int) {
	*a = 20
	fmt.Println(*a * 2)
}

func main() {
	a := 10
	changeVar(&a)
	fmt.Println(a)
}
