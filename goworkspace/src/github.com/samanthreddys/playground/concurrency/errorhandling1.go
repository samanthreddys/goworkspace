package main

import (
	"os"
)

func main() {
	_, err := os.Open("temp.txt")
	if err != nil {
		//fmt.Println("Error Occurred", err)
		//log.Println(err)
		//log.Fatalln(err)
		panic(err)
	}
}
