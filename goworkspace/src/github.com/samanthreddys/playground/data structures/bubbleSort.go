package main

import "fmt"

func bubbleSort(s []int) []int {
	size := len(s) - 1
	for {
		if size == 0 {
			break
		}

		for i := 0; i < size; i++ {
			if s[i] > s[i+1] {
				//fmt.Println("S[i]:", s[i], "i: ", i, "s[i+1]: ", s[i+1])
				tmp := s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
			}
		}
		size--
	}
	return s

}

func main() {
	lst := []int{10, 5, 2, 7, 6, 3, 8, 1}
	sortList := bubbleSort(lst)
	fmt.Println(sortList)
}
