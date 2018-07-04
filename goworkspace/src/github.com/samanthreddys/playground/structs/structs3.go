package main

import "fmt"

//Person struct
type Person struct {
	FName string
	LName string
	Age   int
}

// In below line (p Person) is a reciever for the func
func (p Person) fullName() string {
	return p.FName + " " + p.LName
}

func main() {
	p1 := Person{"Tom", "Doe", 26}
	p2 := Person{"Jane", "Doe", 25}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
