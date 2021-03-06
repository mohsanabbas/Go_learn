package main

import (
	"fmt"
	"net/http"
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
		go checkStatus(l, c)
	}

}

func checkStatus(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might messed up")
		c <- link
		return
	}
	fmt.Println(link, "is good and running with", res.Status)
	c <- link

}
