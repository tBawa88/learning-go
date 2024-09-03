package main

import (
	"fmt"
	"log"
	"net/http"
)

var server = "http://localhost:3000/home"
var urls = []string{server, server, server, server, server}
var sites = []string{"https://google.com", "https://github.com", "https://amazon.com", "https://twitch.tv", "https://youtube.com", "https://www.reddit.com"}
var sites2 = append(sites, sites...) // concatenating 2 slices to make a new slice. We're appending sites with itself by passing all of it's elements as indiviudal values
func main() {

	c := make(chan string)

	for _, url := range sites {
		go checkStatus(url, c)
	}

	// Spawning more go routines
	for url := range c {
		go checkStatus(url, c)
	}

}

func checkStatusNormal(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making the request ", err)
	}

	fmt.Printf("%s  -- %s\n", url, res.Status)
}

func checkStatus(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making the request ", err)
		c <- "Server down"
	}

	c <- url
	fmt.Printf("%s  -- %s\n", url, res.Status)
}
