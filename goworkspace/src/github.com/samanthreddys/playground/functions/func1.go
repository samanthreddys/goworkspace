package main

import (
	"fmt"
)

func main() {
	n := average(10, 2, 3, 4, 5)

	fmt.Println(n)
}

func average(a ...float64) float64 {
	var total float64
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	for _, v := range a {
		total += v
	}
	return total / float64(len(a))
}
