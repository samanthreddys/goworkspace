package main

import (
	"fmt"
	"sort"
)

// People struct
type People []string

func main() {
	studyGroup := People{"Sam", "Tom", "Harry", "Local", "Visu", "Amy"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("sort.StringSlice(studyGroup): ", sort.StringSlice(studyGroup))
	sort.StringSlice(studyGroup).Sort()
	fmt.Println("sort.Reverse(sort.StringSlice(studyGroup)):", sort.Reverse(sort.StringSlice(studyGroup)))
	p := sort.Reverse(sort.StringSlice(studyGroup))
	fmt.Println("&p:", &p)
	sort.Sort(p)
	fmt.Println(studyGroup)

}
