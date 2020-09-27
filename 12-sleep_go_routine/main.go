package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.netflix.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkStatus(link, c)
	}

	// for {
	// 	go checkStatus(<-c, c)
	// }

	// watch the chanle c and it emits value back assign it to variable l
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}

}

func checkStatus(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might fucked up")
		c <- link
		return
	}
	fmt.Println(link, "is good and running with", res.Status)
	c <- link

}
