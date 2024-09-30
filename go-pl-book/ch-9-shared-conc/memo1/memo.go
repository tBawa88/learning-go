package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
*
Designing a concurrent non blocking cache
We want to use caching to **memoize** the result of an expensive function like an http.Get()
*/
type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func              // this could be the function which implements some expensive operation
	cache map[string]result // this is our cache
}

type Func func(key string) (interface{}, error)

// think of this as useMemo()
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Not concurrency safe

/*
	We create an instance of Memo
	Then call this Get method by passing a url string to it
	It checks if the result of the url string is already stored in the map or not
	If not stored, then it calls the 'f' method of memo object
	stores the result into the cache

returns the result
*/
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		// meaning the result is not yet memoized
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
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
