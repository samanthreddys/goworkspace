package main

import "fmt"

func main() {

	n := HashBucket("Sam Reddy", 12)
	fmt.Println(n)

}

// HashBucket fuction will assign a bucket to a word
func HashBucket(word string, buckets int) int {
	fmt.Println(word[0])
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
