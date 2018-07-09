package main

import "fmt"

var lst = []int{4, 9, 2, 7, 1, 3, 6, 8}
var size = len(lst)

func selectionSort(lst []int) []int {
	for i := 0; i < size; i++ {
		var smallest = i
		for j := i; j < size; j++ {
			if lst[j] < lst[smallest] {
				smallest = j
			}
		}
		lst[i], lst[smallest] = lst[smallest], lst[i]
	}
	return lst

}

func main() {
	fmt.Println("size:", size)
	list := selectionSort(lst)
	fmt.Println(list)

}

// Below code is wrapped into a func

/* 	for i := 0; i < size; i++ {

	var smallest = i

	fmt.Println("Smallest inside I loop:", smallest)

	for j := i; j < size; j++ {
		fmt.Println("Inside J List For Each :", lst[j])
		fmt.Println("Inside J List For Each Smallest :", lst[smallest])

		if lst[j] < lst[smallest] {

			fmt.Println("Smallest inside J: ", smallest)
			smallest = j
			fmt.Println("smallest after j assign:", smallest)

		}

	}
	lst[i], lst[smallest] = lst[smallest], lst[i]
}
fmt.Println(lst) */
