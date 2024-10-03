package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Func func(string) (interface{}, error)
// type result struct {
// 	value interface{}
// 	err   error
// }

// type entry struct {
// 	res   result
// 	ready chan struct{} // to broadcast the ready state
// }

// type Memo struct {
// 	F     Func
// 	mu    sync.Mutex // to guard the cache
// 	cache map[string]*entry
// }

// func NewMemo(f Func) *Memo {
// 	return &Memo{F: f, cache: make(map[string]*entry), mu: sync.Mutex{}}
// }

// // this is the concurrent goroutine, will contain all the shared variables and we'll have to implement different methods to make it conc. safe
// func (memo *Memo) Get(url string) (interface{}, error) {
// 	memo.mu.Lock()
// 	ent := memo.cache[url]
// 	if ent == nil {
// 		// entry doesn't exist in the cache, make the request
// 		fmt.Println(url, " entry not found in cache")
// 		ent = &entry{ready: make(chan struct{})}
// 		memo.cache[url] = ent
// 		memo.mu.Unlock() // we release the lock as soon as we make a new entry for this url, now we'll rely on the channel

// 		ent.res.value, ent.res.err = memo.F(url)
// 		close(ent.ready)

// 	} else {
// 		fmt.Println(url, " entry found in cache, waiting on result")
// 		memo.mu.Unlock()
// 		<-ent.ready
// 	}
// 	return ent.res.value, ent.res.err
// }
