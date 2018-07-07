package main

import "fmt"

// constants are parallel type system in GO
const p int = 100

const (
	a  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB
	TB
)
const (
	myName = "Reddy"
	MyID   = "12345"
)

func main() {
	fmt.Println(a)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
