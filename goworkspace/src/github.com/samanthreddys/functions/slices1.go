package main

import "fmt"

func main() {
	mySlice := make([]string, 0, 10)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	for i := 0; i < 1000; i++ {
		mySlice = append(mySlice, string(i))
		//fmt.Println("Lenght:", len(mySlice), "Capacity:", cap(mySlice), "mySlic:", mySlice[i])
	}
	fmt.Println(mySlice)
}
