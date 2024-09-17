package main

import (
	"fmt"
	"net/http"
	"os"
)

// http.HandlerFunc() =====> this is a type that looks like this => type HandlerFunc func(ResponseWriter, *Request)
// it's an adapter type, which let's user define normal functions that match the method signature of it's type, and type cast them into http.Handler
// so that these functions can be provided to a ServerMux as normal http.Handler

func main() {
	db := database{"shoes": 50, "socks": 2.5}
	mux := http.NewServeMux()

	mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/price", db.price) // could also do this, don't have to type cast the function into a handler explicitly

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("Error starting the server ", err)
		os.Exit(1)
	}

}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("%0.2f", d) }

func (db database) list(res http.ResponseWriter, req *http.Request) {
	// render a list of all items in the db
	for item, price := range db {
		fmt.Fprintf(res, "%s:\t%.2f", item, price)
	}
}

func (db database) price(res http.ResponseWriter, req *http.Request) {
	// url.Value is wrapper on map[string] []string
	query := req.URL.Query().Get("item")
	price, ok := db[query]

	if !ok {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "Item doesnot exist in the DB")
		return
	}

	fmt.Fprintf(res, "Price of %s = %.2f", query, price)
}
