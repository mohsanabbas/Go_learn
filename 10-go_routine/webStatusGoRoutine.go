package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://almundo.com.ar",
		"https://ww2.cvc.com.br/",
		"https://www.youtube.com",
		"https://www.blunder.com",
	}

	c := make(chan string) // creating a channel of type string we will pass it down to checkStaus and listen it main func. or recieve back here

	for _, link := range links {
		go checkStatus(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // we are waiting for every go routine to emit one message and then we going to print status
	}
	// fmt.Println(<-c) // recieving messages from channel is blocking code. which ever return first message in channel,
	// main routine going to exit the programm
}

func checkStatus(link string, c chan string) {
	res, err := http.Get(link) // we are sitting here waiting for response for each link get req we will solve this via Go routine and channels
	if err != nil {
		fmt.Println(link, "might messed up")
		c <- "might be down " // here we will emit message back in channel and will listen in main
		return
	}
	fmt.Println(link, "is good and running with", res.Status)
	c <- "Up and running " // here we will emit message back in channel and will listen in main
	// fmt.Println(io.Copy(os.Stdout, res.Body))

}
