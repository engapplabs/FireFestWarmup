package main

import (
	"fmt"
)

//////////////////////////////////////////////////////
type MyError struct {
	message string
}

func NewMyError(message string) *MyError {
	return &MyError{
		message: message,
	}
}

func (myError *MyError) Error() string {
	return fmt.Sprintf("%s", myError.message)
}

/////////////////////////////////////////////////

func divideBy(a, b float64) (float64, error) {
	if b == 0 {
		return -1, NewMyError("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {

	a := 10.0
	b := 0.0

	result, error := divideBy(a, b)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
