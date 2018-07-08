package main

import (
	"fmt"
	"math"
)

// Shape1 Interface
type Shape1 interface {
	area() float64
}

//Circle1 Struct
type Circle1 struct {
	radius float64
}

//Square1 struct
type Square1 struct {
	side float64
}

func (c *Circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square1) area() float64 {
	return s.side * s.side
}

func info1(z Shape1) {
	fmt.Println("area:", z.area())
}

func main() {
	a := Circle1{10}
	b := Square1{5}
	fmt.Println("&a:", &a)
	fmt.Println("&b:", &b)
	info1(&a)
	info1(b)
	info1(&b)
}
