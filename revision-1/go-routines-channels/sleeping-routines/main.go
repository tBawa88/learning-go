package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var sites = []string{"https://google.com", "https://github.com", "https://amazon.com", "https://twitch.tv", "https://youtube.com", "https://www.reddit.com"}

func main() {

	c := make(chan string)

	for _, url := range sites {
		go checkStatus(url, c)
	}

	// Sleeping the go routine before it enters the checkStatus function by using "function literal syntax"
	// A function literal is an anonymous function with no name, and is executed immidiately
	for url := range c {
		go func(url string, c chan string) {
			time.Sleep(time.Second * 1) //when the go routine hits this call, it's put to sleep and the main go routine will move on to next interation and spawn another routine
			checkStatus(url, c)
		}(url, c)
	}

}

func checkStatus(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making the request for ", url)
		c <- "Error"
	}

	fmt.Printf("%s -- %s\n", url, res.Status)
	c <- url
}
