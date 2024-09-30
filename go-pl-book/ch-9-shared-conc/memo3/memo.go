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
	Using 2 locks instead of one.
	First one is to read the cache, then another lock if the record is not found and we need to write new data
	If a goroutine is busy fetching data, others are free to read from the cache

	But this still doesn't solve the problem of 2 URLs being fetched 2 times
	The first goroutine doesn't find, and starts a request. In the meantime, the second goroutine also doesn't find and starts its own result
*/
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)

		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
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
