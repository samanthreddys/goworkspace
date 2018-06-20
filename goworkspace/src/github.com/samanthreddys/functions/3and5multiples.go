package main

import "fmt"

func multiples(x int) int {
	var sum int
	for i := 0; i < x; i++ {
		if (i%5 == 0) || (i%3 == 0) {
			//fmt.Println(i)
			sum += i

		}
	}
	return sum
}

func main() {
	fmt.Println(multiples(1000))
}
