package main

import (
	"fmt"
)

func main() {

	// make(chan int) -- Channel
	// make(chan int,10)-- Buffered Channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Channel Pass:", i)
			c <- i
		}
		close(c) // Channel close -- will stop passing the values-- reciever can still recieve the values
	}()

	for n := range c {
		//fmt.Println("Second Function")
		fmt.Println("Channel Value: ", n)
	}

	//time.Sleep(time.Second)
}
