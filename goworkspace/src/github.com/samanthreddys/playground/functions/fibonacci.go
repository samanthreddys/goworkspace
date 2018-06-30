package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, a+b

		return ret
	}
}

func main() {
	fibo := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}

}
