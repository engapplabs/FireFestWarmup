package main 

import "fmt"

func main() {


	//while like
	i := 1 
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------------------------------")

	//pure for
	for i := 0; i < 10; i++ {
		fmt.Println(i * 2)
	}

	fmt.Println("------------------------------")

	// do while
	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println("OII")
	}	
}
