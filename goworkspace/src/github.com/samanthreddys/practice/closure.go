package practice

import (
	"fmt"
)

var x int = 0

func increment() int {
	x++
	return x
}

func Increment() {
	fmt.Println(increment())
	fmt.Println(increment())
}
