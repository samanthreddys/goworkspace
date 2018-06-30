package main

import "fmt"

func one() {
	fmt.Println("Print Function One!")
}

func two() {
	fmt.Println("Print Function Two!")
}

func main() {
	defer two()
	one()

}
