package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{"http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://instagram.com", "https://viewsourcecode.org/snaptoken/kilo/", "https://www.apple.com/in/"}

func main() {
	c := make(chan string) //made channel of string type

	for _, url := range urls {
		go checkStatus(url, c) //starting the first child routines for each value in slice
	}

	stayChecking(c)

}

func stayChecking(c chan string) {
	for url := range c { //this becomes the blocking call for the main routine now
		go func(url string) {
			time.Sleep(time.Second * 4)
			checkStatus(url, c)
		}(url)
	}
}

func checkStatus(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request to ", url)
		c <- url
		return
	}

	if res.StatusCode != 200 {
		fmt.Println(url, " ", res.Status)
		c <- url
		return
	}

	c <- url
	fmt.Println(url, " ", res.Status)
}
