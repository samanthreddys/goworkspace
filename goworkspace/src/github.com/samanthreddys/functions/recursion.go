package main

import "fmt"

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {
	fmt.Println(factorial(5))
}
