package main

import "fmt"

func main() {
	myMap := make(map[int]string)

	myMap[1] = "India"
	myMap[2] = "UnitedStates"
	myMap[3] = "China"
	myMap[4] = "Australia"

	fmt.Println(myMap)

	for key, val := range myMap {
		fmt.Println(key, ":  ", val)
	}

}
