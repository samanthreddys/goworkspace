package main

import "fmt"

func sum(x *int) {
	fmt.Println(*x)
	*x = 0

}
func main() {
	x := 5
	sum(&x)
	fmt.Println(x)
}
