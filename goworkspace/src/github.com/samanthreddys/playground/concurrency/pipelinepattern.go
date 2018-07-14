package main

import "fmt"

func main() {
	for n := range sq(sq(sq(gen(2, 3)))) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, n := range nums {
			c <- n
		}
		close(c)
	}()
	return c
}

func sq(inp chan int) chan int {
	c := make(chan int)
	go func() {
		for n := range inp {
			c <- n * n
		}
		close(c)
	}()
	return c
}
