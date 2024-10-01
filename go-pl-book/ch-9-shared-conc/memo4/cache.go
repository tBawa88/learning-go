package main

import (
	"fmt"
	"sync"
)

type Func func(string) (interface{}, error)

// Memo is basically a type that we're using to memoize the result of an expensive function call
// - A function who's result we want to memoize
// - It stores a map which acts as the cache and stores the result of the function for instant lookups in future
// - And a mutex to guard this cache from a data race
type Memo struct {
	f  Func
	mu sync.Mutex
	// cache map[string]result
	cache map[string]*entry
}

// Entry struct contains a result from the function call
// an unbuffered channel which we will close once result has been fetched fromt the slow function memo.f
type entry struct {
	res   result
	ready chan struct{}
	// the moment the result is set, this channel will be closed to broadcast to all goroutines
}

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry), mu: sync.Mutex{}}
}

// One problem we're having is redundant writes to the cache. Imagine 2 gorutines trying to look for the same key and don't find their result in the cache
// both of them start their own fetch request and the one who finishes later overwirtes the previous goroutines result in the cache
// We need to preven this (called duplicate supression)

// SOLUTION: Modify the map, instead of directly storing the result in the map, we store another data structure, call it 'entry'
// Each element in the map will be a pointer to this entry struct
// the benefit of storing pointer to entry is we don't have to make another write after data has been fetched for the current key
// since other goroutine will have access to the same memory address
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	ent := memo.cache[key]
	if ent == nil {
		// meaning this is the first request for the key
		//this goroutine is responsible for computing the result and broadcasting the ready state
		ent = &entry{ready: make(chan struct{}, 1)}
		memo.cache[key] = ent
		memo.mu.Unlock()

		// now make the request and set the values
		ent.res.value, ent.res.err = memo.f(key)
		// after this we don't need to update this value in the cache since it's a pointer
		// broadcast ready state
		close(ent.ready)
	} else {
		fmt.Println(key, " found in cache, waiting for ready")
		memo.mu.Unlock()
		<-ent.ready // wait for ready condition
	}
	return ent.res.value, ent.res.err
}

// EXPLAINATION : Imagine the scenario where the entry doesn't exist for the key
// First goroutine will come, create a new entry with a channel and an empty result struct, and push it into the cache
// This goroutine then proceeds to make the http request
// In the meantime, Second goroutine comes for the same key, it will find the the entry does exist in the cache.
// then it proceeds to check if the result is in ready state or not.
// Now if the ready channel has been closed by the first goroutine, then the `<- ent.ready` will resolve immidiately
// This,no 2 goroutine will make the fetch call for the same Key

// also since entry is a pointer, therfore for the same key, all goroutines will be accessing the same entry object
