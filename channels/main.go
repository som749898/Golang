package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, "is working")
	c <- link
	fmt.Println(resp.Status)
}