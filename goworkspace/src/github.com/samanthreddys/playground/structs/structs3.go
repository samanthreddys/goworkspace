package main

import "fmt"

//Person struct
type Person struct {
	FName string
	LName string
	Age   int
}

// Company struct
type Company struct {
	Person
	FName string
	LName string
	State string
}

// In below line (p Person) is a reciever for the func
func (p Person) fullName() string {
	return p.FName + " " + p.LName
}

func (c Company) fullName() string {
	return c.FName + " " + c.LName
}

func main() {
	p1 := Person{"Tom", "Doe", 26}
	p2 := Person{"Jane", "Doe", 25}

	myDetails := Company{
		Person: Person{
			FName: "Tim",
			LName: "Cook",
			Age:   32,
		},
		FName: "John",
		LName: "Doe",
		State: "TX",
	}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(myDetails.fullName())
	fmt.Println(myDetails.Person.fullName())
}
