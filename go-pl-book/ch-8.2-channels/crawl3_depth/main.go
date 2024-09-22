package main

/*
	DEPTH-CONTROLLED LINK CRAWLER

*/

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"tbawa/go-course/go-pl-book/ch-8.2-channels/links"
)

var maxDepth = flag.Int("depth", 3, "Defines the max limit of depth for each link crawled")

var token = make(chan struct{}, 20)

var seenLock = sync.Mutex{}
var seen = make(map[string]bool)

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Depth = %d , Parsing url = %s\n", depth, url)

	if depth >= *maxDepth {
		return
	}

	token <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-token // release a tokenS
	if err != nil {
		log.Print(err)
		return
	}

	// iterate over each link, and recursively spawn a goroutine for each link, while keeping track of the depth
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue // skip this iteration
		}
		seen[link] = true
		seenLock.Unlock() // if it's not seen the lock would still be on
		wg.Add(1)         // since another goroutine is being spawned
		go crawl(link, depth+1, wg)
	}

}

func main() {
	flag.Parse()

	wg := &sync.WaitGroup{} // a pointer to a waitgroupt which can be modified by the crawl function

	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg) // start with 0 depth
	}

	wg.Wait()
}
