package main

import (
	"fmt"

	"github.com/samanthreddys/hello"
	"github.com/samanthreddys/practice"
)

func main() {
	hello.Hello()
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
