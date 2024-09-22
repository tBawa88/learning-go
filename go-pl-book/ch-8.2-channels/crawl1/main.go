package main

import (
	"fmt"
	"log"
	"os"
	"tbawa/go-course/go-pl-book/ch-8.2-channels/links"
)

// First attempt at a concurrent webcrawler, which visits each link provided through the CLI argument
// Uses the built in 'link'

func crawl(url string) []string {
	fmt.Println("Parsing url  = ", url) // printing each URL as it comes
	listOfLinks, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return listOfLinks
}

func main() {
	worklist := make(chan []string) // a channel of string slice

	go func() { worklist <- os.Args[1:] }() // since this is a un-buffered channel, we cannot attempt a send inside the main goroutine
	// there is no worker goroutine listening to this channel therefore the main goroutine would go into a deadlock

	// processing each link concurrently (each iteration in parallel)
	// mainting a map of seen URLs so that we don't call crawl() for the same URL twice
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link) // putting it back into the worklist channel
				}(link)
			}
		}
	}
}

/*
	WHAT WE'RE DOING
	- created a channel that is of type []string, meaning we can only send /receive a string slice
	- created a 'link' package that has an Extract function. This functions takes in a link and sends a GET request
	- obtains the HTML response, parses it using html.Prase(res.Body) package , recursively walks over every node, and extracts the href value of all <a> elements
	- returns those links as  []string, and we're putting that back into the 'worklist' channel
	- the main goroutine keeps spawning more goroutines as long as the data keeps coming into the channel
*/

/*
	PORBLEM WITH THE CODE
	- it's too parallel
	- 2024/09/22 13:39:27 Get "https://www.pcgamesn.com/minecraft/15-best-minecraft-texture-packs":
		dial tcp [2606:4700:10::6816:449a]:443: connect: network is unreachable

	- a DNS lookup failure even the the URL is reliable and should be working
	- this error occurs when there are so many connections at once that the per-process limit of open files is exceeded and calls to net.Dial() are failing
	- this type of error occurs when we allow unbounded concurrency
*/
