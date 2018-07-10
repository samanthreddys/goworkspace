package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var t = time.Now()
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	//var t1 = time.Now()
	var d = time.Since(t)
	fmt.Println(d.Seconds())
}

func foo() {
	for i := 0; i < 10; i++ {
		//fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		//fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
