// Accesing the URL /count returns the number of requests so far, excluding the visits to /count itself
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// It's same as the previous handler, but it's increments the value of count
// But to make sure that 2 parallel requests don't mess up the actual count, it acquires a lock before updating the value of count, and then releases it right after
func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

// echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}
