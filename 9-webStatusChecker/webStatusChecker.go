package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://almundo.com.ar",
		"https://ww2.cvc.com.br/",
	}

	for _, link := range links {
		checkStatus(link)
	}
}

func checkStatus(link string) {
	res, err := http.Get(link) // we are sitting here waiting for response for each link get req we will solve this via Go routine and channels
	if err != nil {
		fmt.Println(link, "might messed up")
		return
	}
	fmt.Println(link, "is good and running with", res.Status)
	// fmt.Println(io.Copy(os.Stdout, res.Body))

}
