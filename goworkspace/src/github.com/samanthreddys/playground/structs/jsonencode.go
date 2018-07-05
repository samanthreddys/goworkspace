package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct
type Person struct {
	Name        string //`json:"Full Name"` // this is used to rename the key name in json
	Age         int
	City        string //`json:"-"` // this is used not to pass the element in json
	notExported bool
}

func main() {
	p1 := Person{"James", 24, "Costa Mesa", false}
	fmt.Println(json.NewEncoder(os.Stdout))
	json.NewEncoder(os.Stdout).Encode(p1)
}
