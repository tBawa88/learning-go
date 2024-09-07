// handler function displayes complete information about the incoming request along with all the headers and method types
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", handler)
	http.HandleFunc("/lisa", handlerLissa)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Error starting the server ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

	// r.Header is a map[string][]string, iterating over the map and printing all of em
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Address = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

func handlerLissa(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}
