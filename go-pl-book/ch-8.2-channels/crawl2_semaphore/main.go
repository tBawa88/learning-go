package main

import (
	"fmt"
	"log"
	"os"
	"tbawa/go-course/go-pl-book/ch-8.2-channels/links"
)

/*
	SOLUTION: (to crawl1)
	- Limit the number of parallel access to a resource to match the level of parallelism that is available
	- in this case, we can make sure that no more than 'n' calls to links.Extract() are active at any given time
	- 'n' being the limit for open file descriptors (which are created when we open a file)

	- This way of limiting the concurrency, using a buffered channel is called "counting semaphore"
	- Each of the 'n' vacant slots in the buffer represents 'n' available tokens which the holder(goroutine) can use to proceed
	- When a value is sent to the channel, a token is used
	- When a value is received from the channel, a token is released. Creating a new vacant slot

	- So in order to make a call to link.Extract(), the goroutine must first try to send value to the channel
	- if there is space in the buffer, the goroutine can go ahead, otherwise it must wait untill another goroutine receives from  the channel (indicated it's finished task)
*/

// Declaring a global buffered channel in which empty places represent a token
var token = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println("Parsing url ", url)

	token <- struct{}{} // acquired a token, by occupying one seat. If the buffer is full, this routine will have to wait untill a seat is empty
	list, err := links.Extract(url)
	<-token // release the acquired token after making functoin call for this url

	if err != nil {
		log.Print(err)
		return nil
	}
	return list
}

/*
Another thing we need to fix is that the main goroutine is never exiting since it's always listening to the worklist inside the for loop
For the program to exit, the main loop needs to be terminated when the worklist is empty and ther are no active gorutines
This can be fixed using a simple counter integer which keeps track
Instead of running the loop by constantly receiving from the channel, we run the loop using a counter variable
Each time a goroutine is spawned, the coutner is upped. And each time all links received from worklist have been iterated over,
counter is decremented (automatically by loop)
*/
func main() {
	worklist := make(chan []string)
	var n int
	go func() { worklist <- os.Args[1:] }()

	n++ // first for the os.Args, since that must be processed
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		linkList := <-worklist // receive from the channel and start a loop
		for _, link := range linkList {
			if !seen[link] {
				seen[link] = true
				n++ // if new link is found, we increment 'n' since a new send will be made to worklist, and we need another iteration to receive that
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
