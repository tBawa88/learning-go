package main

// http.Handler inteface
// For a type to qualify as http.Handler, it must implement a ServerHTTP(ResponseWriter, *Request)
// To start a server, we call ListenAndServer(string, http.Handler) function. it takes in an address and calls the ServerHTTP() method of the Handler
// Think of the handler as the const app = express(). IT handles all the requests and route them to their correct route handlers

// In this we're implementing our own http.Handler
// The database type will act as a Handler but when it's ServerHTTP function is called,
import (
	"fmt"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 2.5}

	http.ListenAndServe("localhost:8080", db)
}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("%0.2f", d) }

func (db database) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	for key, value := range db {
		fmt.Println(key, value)
	}
}

// when the server starts, it prints all key values into the Stdout
