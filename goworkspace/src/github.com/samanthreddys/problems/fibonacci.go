package main

import "fmt"

func main() {
	var n int
	n = 10
	for i := 0; i <= n; i++ {
		fmt.Println(fib(i))
	}

}

func fib(n int) int {
	//var b int

	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
	//fmt.Println("Fibo Sequence:", b)

}
