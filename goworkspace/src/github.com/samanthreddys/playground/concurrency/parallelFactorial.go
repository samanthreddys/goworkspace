package main

import "fmt"

func main() {
	in := gen()
	f := fact(in)
	for n := range f {
		fmt.Println(n)
	}

}
func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 15; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(inp <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range inp {
			//fmt.Println("N: ", n)
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func factorial(f int) int {

	total := 1
	for i := f; i > 0; i-- {
		total *= i
		//fmt.Println("Total:", total)
	}
	return total
}
