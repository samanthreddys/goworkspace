package main

import "fmt"

func main() {
	x := 10
	y := 20

	x = x + y
	y = x - y
	x = x - y
	fmt.Println("x,y", x, y)
}
