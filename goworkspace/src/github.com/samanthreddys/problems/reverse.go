package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 10; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i

		}
	}
	fmt.Println("Print i:", sum)
}
