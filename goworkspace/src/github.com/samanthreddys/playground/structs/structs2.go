package main

import "fmt"

type employee struct {
	ID     string
	Name   string
	Salary float64
}

type company struct {
	employee
	City  string
	State string
	Zip   int
}

func main() {

	myDetails := company{
		employee: employee{
			ID:     "12345",
			Name:   "Tom",
			Salary: 456.421,
		},
		//company.ID: "CX2780",
		City:  "Plano",
		State: "TX",
		Zip:   75024,
	}

	fmt.Println(myDetails)

}
