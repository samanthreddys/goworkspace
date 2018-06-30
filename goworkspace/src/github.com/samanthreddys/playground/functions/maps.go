package main

import "fmt"

func main() {

	myMap := make(map[string]int)
	myMap["a1"] = 7
	myMap["a2"] = 8
	myMap["b1"] = 9
	myMap["b2"] = 10

	fmt.Println(myMap)
	fmt.Println(len(myMap))
	v1 := myMap["b1"]
	fmt.Println(v1)
	delete(myMap, "b1")
	_, ok := myMap["b1"]
	fmt.Println(ok)

	fmt.Println("")

}
