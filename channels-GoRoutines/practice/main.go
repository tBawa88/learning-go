package main

import (
	"fmt"
	"net/http"
	"time"
)

// a program that keeps pinging my website to keep it form sleeping

func main() {
	url := "https://job-tracker-73fh.onrender.com/login"

	ch := make(chan string)

	go pingSite(url, ch) //just make sure to get things rolling by starting a go routine
	for status := range ch {
		go func() {
			time.Sleep(time.Second * 5)
			fmt.Println("status ", status)
			pingSite(url, ch)
		}()

	}

}

func pingSite(url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request")
		ch <- "Exit"
		return
	}

	fmt.Println(res.Status)
	ch <- res.Status
}
