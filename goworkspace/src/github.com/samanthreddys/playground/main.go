package main

import (
	"fmt"

	"github.com/samanthreddys/playground/Hello"
	"github.com/samanthreddys/playground/practice"
)

func main() {
	Hello.Hello()
	fmt.Println(practice.MyName)
	fmt.Println(practice.myCountry)
	practice.MyVariables()
	fmt.Println(practice.MyTotal())
	practice.Increment()

	x := 0
	increment := func() int {
		x++
		return x

	}
	fmt.Println(increment())
	fmt.Println(increment())

}
