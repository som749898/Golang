package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links {
		go getWebsite(link, c)
	}
	for l := range c {
		// go getWebsite(l, c)
		go (func() {
			time.Sleep(5 * time.Second)
			getWebsite(l, c)
		})()
	}
}

func getWebsite(l string, c chan string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println("Website not working", l)
		c <- l
		return
	}
	fmt.Println(l)
	fmt.Println("Website working fine", resp.Status)
	c <- l
}
