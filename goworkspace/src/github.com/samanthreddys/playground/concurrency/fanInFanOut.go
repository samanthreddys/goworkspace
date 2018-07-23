package main

import (
	"fmt"
	"sync"
)

func main() {
	g := gen(2, 3)
	c1 := sq(g)
	c2 := sq(g)
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(n ...int) chan int {

	fmt.Printf("Type of Numbers %T\n", n)

	out := make(chan int)
	go func() {
		for _, i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(c chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range c {
			out <- i * i
		}
		close(out)
	}()

	return out
}

func merge(c ...chan int) chan int {

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, n := range c {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()

		}(n)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
