```go
    package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{"http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://instagram.com", "https://viewsourcecode.org/snaptoken/kilo/", "https://www.apple.com/in/"}

func main() {

	//treat a channel just like any other value of struct or a map. We can pass it around our program
	c := make(chan string)
	for _, url := range urls {
		//for this channel to be used by this child routine for comunnication with the main routine, we must pass it to the function that is being executed by the child routine
		//so that it can be used inside that function
		go checkStatus(url, c)
	}
	// fmt.Println(<-c) //waiting for the value to come out of channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)		//this channel will never get a value passed to it, because there is no child routine that is still going to be running when the main routine comes here

	//listening to the channel exactly len(slice) number of times
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	//for repeated pinging of these sites
	// for {
	// 	go checkStatus(<-c, c)
	// }

	//improving for loop syntax, can use range keyword to extract the values coming out of this channel
	// for url := range c {
	// 	go checkStatus(url, c)
	// }

	//sleeping a routine to create a delay between requests
	//time.Sleep() is a special function in the time package, which is used to sleep a routine for specified amount of time
	for url := range c {
		//placing this sleep function here will block the main routine and it will only be able to listen the data from the channel, once every 2 seconds
		//this could throttle the channel as the children routines finish up and send their messages into the channel before the main routine is ready to listen to those messages
		// time.Sleep(time.Second * 2)
		// go checkStatus(url, c)

		//using a function literal (equivalent to anonymous function in JS)
		//one unspoken rule is that never share a variable between the main routine and child routines ever, always create a copy, and pass it to the child routine
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkStatus(url, c)
		}(url) //invoking this function immidiately
	}

}

// even here we must specify what type of data we wish to share inside the channel
func checkStatus(url string, c chan string) {
	//sleeping the child routine is also not completely appropriate, since we expect checkStatus() to immidiately check the status for us without any pause
	// time.Sleep(time.Second * 5)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request to ", url)
		c <- url
		return
	}
	if res.StatusCode != 200 {
		fmt.Println(url, " is down ", res.Status)
		c <- url
		return
	}
	fmt.Println(url, " ", res.Status)
	c <- url
}

```