package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 10)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	for i := 0; i < 180; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Lenght:", len(mySlice), "Capacity:", cap(mySlice), "mySlic:", mySlice[i])
	}
}
