package practice

import "fmt"

var MyName = "samanth reddy"
var myCountry = "India"

// short hand variables can be used only inside func
// zero value- if you dont explicitly define for a variable type it will be set as zero
//%T for type format
//  var and shorthand preferred methods to declare variables. Not preferredmethods are assign declare and initialize

// shorthand declare
func MyVariables() {
	a := 10
	b := true
	c := "golang"
	var d float64
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)

}
