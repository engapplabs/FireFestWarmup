package main

import (
	"fmt"
)

func main() {

	var s = make([]string, 3)
	fmt.Println(s)

	s[0] = "vaca"

	s[1] = "boi"

	s[2] = "bode"

	fmt.Println(s)

	s = append(s, "oi")
	fmt.Println(s)

	//copying

	var second = make([]string, len(s))
	copy(second, s)
	fmt.Println(second)

	// initizalize and declare
	var x = []string{"Aurelio", "Leandro"}
	fmt.Println(x)

}
