package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person struct
type Person struct {
	Name        string //`json:"Full Name"` // this is used to rename the key name in json
	Age         int
	City        string //`json:"-"` // this is used not to pass the element in json
	notExported bool
}

func main() {
	var p1 Person
	es := strings.NewReader(`{"Name":"James", "Age":24, "City":"Costa Mesa", "notExported":false}`)
	json.NewDecoder(es).Decode(&p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.City)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1)

}
