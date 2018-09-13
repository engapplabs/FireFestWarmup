package main

import (
	"fmt"
)

func getTimes2And3(number int) (int, int) {
	return number * 2, number * 3
}

func main() {
	a, b := getTimes2And3(2)
	fmt.Println(a, b)
}
