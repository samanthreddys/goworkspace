package main

import (
	"fmt"
	"time"
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
	}()

	go func() {
		for {
			fmt.Println("Second Function")
			fmt.Println("Channel Release: ", <-c)
		}
	}()
	time.Sleep(time.Second)
}
