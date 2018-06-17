package main

import (
	"fmt"
)

func main() {
	x := 51 % 10
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	var m int
	fmt.Println("Enter a number")

	fmt.Scan(&m)
	for i := 0; i <= m; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
	}

}
