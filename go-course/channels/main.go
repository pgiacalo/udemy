package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.donaldjtrump.com/",
		"https://www.donaldjtrump.com/",
		"https://www.donaldjtrump.com/",
		"https://www.donaldjtrump.com/",
		// "https://www.google.com",
		// "https://www.facebook.com",
		// "http://www.stackoverflow.com",
		// "https://www.golang.org",
		// "https://www.amazon.com",
	}

	c := make(chan string)
	var count int = 0

	for _, link := range links {
		go checkLink(link, c, count)
	}

	for l := range c {
		count++
		go func(link string) {
			//			time.Sleep(1 * time.Millisecond)
			checkLink(link, c, count)
		}(l)
	}
}

func checkLink(link string, c chan string, count int) {
	time.Sleep(time.Second)
	_, err := http.Get(link)
	if err != nil {
		c <- link
		fmt.Println(link, "is down", count)
		return
	}

	fmt.Println(link, "is up", count)
	c <- link
}
