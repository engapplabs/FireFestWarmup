package main

import (
	"fmt"
)

func getProperFunction(parameter int) func() int {
	var functionToReturn func() int
	switch parameter {
	case 1:
		functionToReturn = func() int {
			return parameter
		}
	case 2:
		functionToReturn = func() int {
			return parameter * 2
		}
	case 3:
		functionToReturn = func() int {
			return parameter * 3
		}
	}
	return functionToReturn
}

func main() {

	var functionHandler = getProperFunction(2)
	fmt.Println(functionHandler())
	fmt.Println(functionHandler)

	var r1 = getProperFunction(1)()
	fmt.Println(r1)
	var r2 = getProperFunction(2)()
	fmt.Println(r2)
	var r3 = getProperFunction(3)()
	fmt.Println(r3)

}
