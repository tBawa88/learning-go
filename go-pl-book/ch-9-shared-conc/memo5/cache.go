package main

// Problem with memo4 : the map is a shared variable even though it's being protected by locks
// ent.res.value and ent.res.err are also being read by multiple goroutines but they're still safe from a Data race, why?
// because the when the write is happening to res.value and res.err, no other goroutine can access them untill the writer goroutine has closed the channel
// But still we can make a monitor goroutine that will confine them and other goroutines will communicate via messages

type Func func(string) (interface{}, error)
type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result // a write only channel which pushes a result object to the channel
}

// modified Memo type
type Memo struct{ requests chan request }

func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e

			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready         // wait for ready
	response <- e.res // send the response to channel
}
