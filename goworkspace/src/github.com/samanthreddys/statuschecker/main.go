package main

import (
	"fmt"
	"net/http"
)

func main() {
	websiteUrl := []string{"http://google.com", "http://localhost:8080/prweb/PRServlet", "http://facebook.com", "http://golang.org", ""}

	c:= make(chan string)
	for _, url := range websiteUrl {
		go statusCheck(url,c)
		
		

	}
	fmt.Println(<-c)
	
}

func statusCheck(url string, c chan string) string {
	resp, err := http.Get(url)
	if err != nil {
		c<-"Error reaching website"
		fmt.Println(url, " is not responding:", err)
		return "Error in reaching website"
	}
	fmt.Println(url + ":" + string(resp.Status))
	c<-string(resp.Status)
	return url + ":" + string(resp.Status)
}
