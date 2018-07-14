package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

	m := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			m <- i
		}
		close(m)
	}()
	for n := range m {
		fmt.Println(n)
	}
}
