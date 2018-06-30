package main

import (
	"fmt"
)

func main() {
	transactions := make([][][]int, 0)

	for i := 0; i < 4; i++ {
		transaction := make([][]int, 0)

		for j := 0; j < 4; j++ {
			//transaction = append(transaction, j)
			innerTransaction := make([]int, 0)
			for k := 0; k < 4; k++ {
				innerTransaction = append(innerTransaction, k)
				//fmt.Println("\n")
			}
			transaction = append(transaction, innerTransaction)
			//fmt.Println("\n")

		}
		//fmt.Println("\n")
		transactions = append(transactions, transaction)

	}
	transactions1 := make([][]int, 0)
	for x := 0; x < 5; x++ {
		transaction1 := make([]int, 0)

		for y := 0; y < 4; y++ {
			transaction1 = append(transaction1, y)

		}
		transactions1 = append(transactions1, transaction1)
	}
	fmt.Println("2D Array:", transactions1)
	fmt.Println("3D Array:", transactions)

}
