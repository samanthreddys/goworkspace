package main
import "fmt"

func filter(numbers []int, callback func(int) bool) []int{
	xs:=[]int{}
	for _, n:= range numbers{
		//fmt.Println("n value is:",n)
		//x :=n%2==0
		if callback(n) {
			if n%2  == 0{
				xs = append(xs,n)

			}
			
		}
	}
	return xs
}

func main(){
	xs:= filter([]int{1,2,3,4,5,6,16,17,7,8,8,14,9,10}, func(n int) bool{
		return n>0
	})
	fmt.Println(xs)
}