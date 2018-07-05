package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.City)
	bs := []byte(`{"Name":"Harry","Age":25,"City":"Costa Mesa"}`)

	json.Unmarshal(bs, &p1)
	fmt.Printf("%T \n", p1)

	fmt.Println("&:", &p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.City)
	fmt.Println(p1)

}
