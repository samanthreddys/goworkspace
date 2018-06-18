package main
import "fmt"

func filter(numbers []int, callback func(int) bool) []int{
	xs:=[]int{}
	fmt.Println("n value is:,b")
	for _, n:= range numbers{
		fmt.Println("n value is:,c")

		if callback(n) {
			fmt.Println("n value is:,d")
			xs = append(xs,n)
			
		}
	}
	fmt.Println("n value is:,f")
	return xs
	
}

func main(){
	xs:= filter([]int{1,2,3,4,5,6,16,17,7,8,8,14,9,10}, func(n int) bool{
		fmt.Println("n value is:,a")
		return n %2 !=0
		
	})
	fmt.Println(xs)
}