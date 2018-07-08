package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6
	fmt.Println(x)
	z, _ := strconv.Atoi(x)
	fmt.Println(z + y)
}
