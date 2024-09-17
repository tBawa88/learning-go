package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// The task is to create a dummy CRUD REST- API which operated on this DB
// One thing to keep in mind is this introduces ""shared variable concurrency""
// Meaning more than 1 requests can mutate the database state and it can produce corrupted data

func main() {
	db := database{"shoes": 50, "socks": 2.5}
	mux := http.NewServeMux()

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("Error starting the server ", err)
		os.Exit(1)
	}

}

type dollars float32

var dblock sync.Mutex

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

func (db database) create(res http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)

	dblock.Lock()
	db[item] = dollars(price)
	dblock.Unlock()

	fmt.Fprintf(res, "Data created, go to /list to check")
}
