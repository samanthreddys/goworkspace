package main

import "fmt"

func linearSearch(lst []int, number int) int {

	for i := 0; i < len(lst); i++ {

		if lst[i] == number {

			return i
		}

	}
	return -1
}

func main() {
	lst := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	var number int
	fmt.Println("Enter a number for lookup:")
	fmt.Scan(&number)
	fmt.Println(linearSearch(lst, number))

}
