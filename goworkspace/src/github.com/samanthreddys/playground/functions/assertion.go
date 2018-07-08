package main

import "fmt"

func main() {
	var x interface{} = "sydney"
	fmt.Println(x)

	key, value := x.(string)

	if value {
		fmt.Printf("%T \n", key)
		fmt.Println(key)
	} else {
		fmt.Println("Key is not a string")
	}

	var number interface{} = 7
	fmt.Printf("%T \n", number)
	fmt.Println(number.(int) + 6)

	// conversion example

	var val1 = 7.6345
	fmt.Printf("%T", val1)
	fmt.Println("%T \n", int(val1))

	// assertion example

	var val2 interface{} = 7.2345
	fmt.Printf("%T \n", val2)
	fmt.Println("%T \n", val2.(float64))

}
