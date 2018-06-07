package main

import "fmt"

const mtsToYds float64 = 1.09361

func main() {
	var mts float64
	fmt.Print("Enter the number of meters: ")
	fmt.Scan(&mts)
	yards := mts * mtsToYds
	fmt.Println("Number of yards: ", yards)

}
