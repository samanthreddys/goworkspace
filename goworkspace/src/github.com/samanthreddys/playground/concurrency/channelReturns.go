package main

import "fmt"

func main() {
	c := Incrementor()
	puller := Puller(c)
	fmt.Println(puller)
	for n := range puller {
		fmt.Println(n)
	}
}

//Incrementor function
func Incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// Puller function
func Puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {

			sum += n
			fmt.Println(sum)
		}
		out <- sum
		close(out)

	}()
	return out
}
