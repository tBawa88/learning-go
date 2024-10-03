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

// A request is a message , it contains a key which will be eventually passed to Func function
// it contains a response channel to which we can send the result of Func function
type request struct {
	key          string
	responseChan chan<- result // a write only channel which pushes a result object to the channel
}

// Memo struct now only contains a channel to which request messages are pushed ( a channel of channles )
// Memo is not a cache anymore, it's just a middleman which listens to a request and forwards it to the monitor goroutine(server)
type Memo struct{ requests chan request }

// It returns a *Memo, and starts another goroutine 'server' which is our MONITOR goroutine (confines the map)
func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get is a goroutine that gets a key, then makes a response channel (of type result)
// constructs a new request message, and passes it to the 'requests' channel of Memo
// after that it waits and listens for a result to come out of IT'S own response channel
// this response will be sent by the deliver() goroutine who is a child of server() go routine
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

// server() goroutine is started whenever a new Memo object is created
// it takes a function and waits for any request message to come out of memo.requests channel
// when a request comes, it fist checks whether the request.key is present in the map or not
// if not, it creates a new entry with a broadcast channel called 'ready'
// then calls e.call() goroutine
// and finally calls e.deliver() goroutine
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e

			go e.call(f, req.key)
		}
		go e.deliver(req.responseChan)
	}
}

// call() goroutine calls the 'f Func', by passing the key
// populates the result field in the entry and closes the 'ready' channel
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

// deliver() goroutine waits for the ready channel to close
// and send the result into the response channel of 'current entry'
// From here the result reaches to Get() goroutine which is waiting for a response to come out of it's response channel
func (e *entry) deliver(responseChan chan<- result) {
	<-e.ready             // wait for ready
	responseChan <- e.res // send the response to channel
}
