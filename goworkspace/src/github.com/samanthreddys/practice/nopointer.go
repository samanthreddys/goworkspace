package main

import "fmt"

func sum(x int) {
	fmt.Printf("%p \n", &x)
	//fmt.Println(&x)
	x = 0

}
func main() {
	x := 20
	fmt.Printf("%p \n", &x)
	//fmt.Println(&x)
	sum(x)
	fmt.Println(x)

}
