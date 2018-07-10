package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func main() {
	var t = time.Now()
	wg.Add(2)
	go incrementor("Foo:")

	go incrementor("Bar:")

	wg.Wait()
	fmt.Println("Final Counter:", counter)
	//var t1 = time.Now()
	var d = time.Since(t)
	fmt.Println(d.Seconds())
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(3 * time.Millisecond))
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}
