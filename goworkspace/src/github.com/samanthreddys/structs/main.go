package main

import "fmt"

type address struct {
	addLine1 string
	addLine2 string
	Zip      int
}

type person struct {
	firstName string
	lastName  string
	contact   address
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson", contact: address{addLine1: "Oen", addLine2: "wer", Zip: 1234}}

	alex.updateName("Hello")
	alex.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
