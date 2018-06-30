package main

import (
	"fmt"
)

func main() {
	data := []float64{10, 203, 30, 40, 50, 70, 90}
	n := average(data...)

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
