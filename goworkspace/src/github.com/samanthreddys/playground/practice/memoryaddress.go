package main

import "fmt"

var b string = "sam"

func main() {
	a := 43
	fmt.Println("Memory address of a is :", &a)
	fmt.Println("Memory address of b is :", &b)
	fmt.Printf("%d \n", &a)
	fmt.Printf("%d \n", &b)

}
