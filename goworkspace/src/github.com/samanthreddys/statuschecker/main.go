package main

import (
	"fmt"
	"net/http"
	"time"
	//	"time"
)

func main() {
	websiteUrl := []string{"http://google.com", "http://localhost:8080/prweb/PRServlet", "http://facebook.com", "http://seek.com.au"}

	c := make(chan string)
	for _, url := range websiteUrl {
		go statusCheck(url, c)

	}
	for l := range c {
		go func(l string) {
			//	fmt.Println("Inside l range for loop:", l)
			time.Sleep(5 * time.Second)
			go statusCheck(l, c)

		}(l)
	}

}

func statusCheck(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in accessing url:", url)
		c <- url
		return
	}
	fmt.Println("URL is active:", url)
	c <- url

}
