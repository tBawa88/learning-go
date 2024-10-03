package main

import "fmt"

// we know there are 2 patterns of making our programs concurrency safe. One is to use mutual exclusion to protect shared resources
// other is to confine all shared resources into one monitor goroutine and make other goroutines communicate with this goroutine by sending messages
// which is also called the principle of 'don't communicate by sharing memory, share memory by communicating'
// each goroutine gets it's own channel through which it communicates with the monitor
// the monitor, for the sake of being fast, spawns other goroutines to handle expensive tasks
// the client goroutine wait for a response to come out of their channels

type Func func(string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}
type entry struct {
	res   result
	ready chan struct{}
}

// Every goroutine will create a request object and send it through the channel to monitor goroutine
type request struct {
	key          string        // url of every request
	responseChan chan<- result // where to send the result for this url
}

// Memo will now only store a channel of channel through which goroutines will communicate with monitor goroutine
type Memo struct {
	requests chan request
}

// this function will start the monitor goroutine by passing it the get function
func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// this is the function called by the main goroutine for every url request
func (memo *Memo) Get(url string) (interface{}, error) {
	responseChan := make(chan result)
	request := request{key: url, responseChan: responseChan}
	memo.requests <- request

	response := <-responseChan // after creating a response channel which is local to this gorutine, we send it to requests channel of memo object
	return response.value, response.err
}

// monitor goroutine, maintains a local cache, which is confined to it
// for every request that comes out of this channel, it spawns 2 more goroutine to handle the blocking calls so that it can remain free to handle more requests
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	// This loop is very important in sequential processing of requests
	// The server goroutine takes out one request out from the channel and processes it
	// Even if 2 duplicate URLs are requested at the same time, since the channel is unbuffered, one of the requests will have to wait untill the previous iteration of thelooop
	// has finished. And when the time comes to process the duplicate requests, the entry will already be present inside the cache(due to the previous request)
	// and all it has to do is wait for the ready channel to close
	for req := range memo.requests {
		ent, ok := cache[req.key]
		if !ok {
			fmt.Println("For ", req.key, " entry not found")
		} else {
			fmt.Println("For ", req.key, " entry found")

		}

		if ent == nil {
			ent = &entry{ready: make(chan struct{})}
			cache[req.key] = ent

			// make the http request, and then close the channel
			go ent.sendRequest(f, req.key)
		}
		// if the ready has been closed, deliver back the response to this req channel
		go ent.deliver(req.responseChan)
	}
}

func (memo *Memo) Close() { close(memo.requests) }

func (ent *entry) sendRequest(f Func, key string) {
	ent.res.value, ent.res.err = f(key)
	close(ent.ready)
}

func (ent *entry) deliver(response chan<- result) {
	<-ent.ready         // wait for this channel to close
	response <- ent.res // send the result instance on 'ent' to this channel
}
