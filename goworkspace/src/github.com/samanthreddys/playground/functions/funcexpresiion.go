package main

import "fmt"

func main() {
	greeting := func() string {
		return "Hello Samanth"
	}
	greeting()
	fmt.Printf(" %T \n", greeting)
	fmt.Println(greeting())
}
