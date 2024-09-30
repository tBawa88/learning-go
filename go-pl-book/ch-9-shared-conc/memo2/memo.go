package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func // this could be the function which implements some expensive operation
	mu    sync.Mutex
	cache map[string]result // this is our cache

}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result), mu: sync.Mutex{}}
}

// Now this is concurrency safe
/*
	But this slows down the concurrency and makes it sequential
	Because each goroutine must wait in line to access the cache
*/
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		// meaning the result is not yet memoized
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

/*
Example of how this might work
*/
func main() {
	memo := New(httpGetBody)
	wg := sync.WaitGroup{}
	for _, url := range incomingURL() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := memo.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s , %s , %.3d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()

}

func incomingURL() []string {
	return []string{"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io", "https://golang.org"}
}

// let'ssay incomingURL is some function that is sending out a stream of URLs
//the whole point of this was we don't want to send the Get request for the same URL twice since it's expensive
