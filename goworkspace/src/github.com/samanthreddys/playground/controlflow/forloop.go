package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	j := 0
	for {
		j++
		if j%2 == 0 {

			continue

		}
		fmt.Println(j)
		if j >= 50 {

			break
		}

	}

	fmt.Println(`samanth reddy`)
}
