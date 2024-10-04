package main

import (
	"context"
	"fmt"
	"net/http"
)

type KeyType string

// One of the use case of context is to pass request specific values to other parts of code
func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// create a context with a key-value
		ctx := context.WithValue(req.Context(), KeyType("userId"), 12345)

		// pass this context as the first argument to another process
		processRequest(ctx)
	})

	http.ListenAndServe(":8080", nil)
}

func processRequest(ctx context.Context) {
	userId := ctx.Value(KeyType("userId"))
	fmt.Printf("User id of current request = %d\n", userId.(int))
}
