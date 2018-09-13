package main

import "fmt"

func main() {

	var myMap = make(map[int]string)

	myMap[0] = "Aurelio"

	myMap[1] = "vaca"

	fmt.Println(myMap)

	delete(myMap, 0)

	fmt.Println(myMap)

}
