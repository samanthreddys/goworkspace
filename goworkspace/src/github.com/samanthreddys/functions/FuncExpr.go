package main

import "fmt"

func main() {
	greet := greeting()
	fmt.Println(greet())
}
func greeting() func() string {
	return func() string {
		return "Hello Sam!"
	}
}
