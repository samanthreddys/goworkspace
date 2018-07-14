package main

import "fmt"

func main() {
	f := factorial(5)
	for n := range f {
		fmt.Println("N Value: ", n)
	}

}
func factorial(f int) chan int {
	c := make(chan int)
	total := 1
	go func() {
		for i := f; i > 0; i-- {
			total *= i
			//fmt.Println("Total:", total)
		}
		c <- total
		close(c)

	}()
	return c
}
