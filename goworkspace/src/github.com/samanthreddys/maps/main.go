package main

import "fmt"

func main() {

	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"Red":    "#ff0000",
		"White":  "#ff5678",
		"Yellow": "#ff0006",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
