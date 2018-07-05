package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	Name        string `json:"Full Name"` // this is used to rename the key name in json
	Age         int
	City        string `json:"-"` // this is used not to pass the element in json
	notExported bool
}

func main() {

	p1 := Person{"Harry", 25, "Costa Mesa", true}

	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)

	fmt.Println(string(bs))
}
