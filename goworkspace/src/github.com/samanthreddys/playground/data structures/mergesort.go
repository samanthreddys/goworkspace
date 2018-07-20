package main

import "fmt"

func main() {
	lst := []int{3, 5, 9, 4, 2, 7, 6, 1, 8, 0}
	fmt.Println("unsorted:", lst)
	fmt.Println(MergeSort(lst))

}

//MergeSort function
func MergeSort(lst []int) []int {
	//fmt.Println("inside merge sort")

	size := len(lst)
	if size < 2 {
		return lst
	}
	mid := size / 2
	var (
		left  = make([]int, mid)
		right = make([]int, size-mid)
	)
	for i := 0; i < size; i++ {
		if i < mid {
			left[i] = lst[i]
		} else {
			right[i-mid] = lst[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))

}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
