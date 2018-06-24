package main

import "fmt"

var myList []string

func main() {

	myList = []string{"s", "a", "m", "a", "n", "t", "h"}
	myslice := []string{"s", "a", "m", "a", "n", "t", "h"}

	fmt.Println(myslice)
	fmt.Println(myList)
	fmt.Printf("% T\n", myslice)
	fmt.Println(myslice[2])
	fmt.Println(myslice[2:5])
	fmt.Println("Reddy"[2])
}
