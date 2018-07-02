package main

import "fmt"

// struct is a aggregate type
func main() {
	type Person struct {
		Name string
		City string
		ID   int
	}
	type age int

	p1 := Person{"Tom", "California", 123}
	p2 := Person{"Harry", "Las Vegas", 145}
	fmt.Println(p1.Name, p1.City, p1.ID)
	fmt.Println(p2.Name, p2.City, p2.ID)

	var myAge age
	myAge = 45

	fmt.Printf("%T %v \n ", myAge, myAge)

}
