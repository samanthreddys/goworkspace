package main

import (
	"fmt"
	"math"
)

// Square struct
type Square struct {
	side float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Shape interface
type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side

}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {

	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{5}
	c := Circle{10}
	//fmt.Println("Hello")
	info(s)
	info(c)
}
