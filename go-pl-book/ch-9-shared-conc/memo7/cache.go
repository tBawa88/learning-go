package main

type Func func(string) (interface{}, error)

type Result struct {
	value interface{}
	err   error
}

type Entry struct {
	res   Result
	ready chan struct{}
}

type Request struct {
	key          string      // requeted resource
	responseChan chan Result // addrr of the client, this is where the Result will be sent
}

// instead of storing the Func as a field, this time we're starting a new goroutine called server which will memoize this function
// cache is also moved to that function so that it's confined to a single goroutine
type Memo struct {
	requests chan Request
}

func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan Request)}
	go memo.server(f) // since server needs access to this requests channel
	return memo
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) Get(url string) (interface{}, error) {
	responseChan := make(chan Result)
	req := Request{key: url, responseChan: responseChan}
	memo.requests <- req
	res := <-responseChan
	return res.value, res.err
}

// monitor goroutine, maintains a local cache which is not accessible by any other goroutine
// it uses the 'requests' channel on Memo object and "LISTENS TO EACH REQUEST SEQUENTIALLY"
// this is important to remove duplicate calls
// if the entry is not found in cache, a new entry is created and a new goroutine is called to handle http call which is expensive
// regardless of whether the entry is found or not, the control reaches to another goroutine called deliver
// this routinne is waits for the 'ready' channel to close, before sending a response back to the responseChan
func (memo *Memo) server(f Func) {
	cache := make(map[string]*Entry)
	for req := range memo.requests {
		ent := cache[req.key]
		if ent == nil {
			ent = &Entry{ready: make(chan struct{})}
			cache[req.key] = ent

			go ent.send(req.key, f)
		}
		go ent.deliver(req.responseChan)
	}
}

func (entry *Entry) send(key string, f Func) {
	entry.res.value, entry.res.err = f(key)
	close(entry.ready)
}

func (entry *Entry) deliver(response chan Result) {
	<-entry.ready // once the ready channel has been closed, meaning entry has now result value
	response <- entry.res
}
