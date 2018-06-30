package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map1[1] = "samanth"
	map1[2] = "chaitu"
	map1[3] = "reddy"

	if val, exists := map1[2]; exists {
		delete(map1, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("The value does not exists")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}
	fmt.Println(map1)
}
