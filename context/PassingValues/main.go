package main

import (
	"context"
	"fmt"
)

type KeyType string

func main() {
	key := KeyType("userId")
	// context.WithValue() doesn't return a CancelFunc since it's not used for running any operation. It's only there for holding values
	// request - scoped vlaues (whatever that is)
	ctx := context.WithValue(context.Background(), key, 123456)

	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	userId := ctx.Value(KeyType("userId"))
	fmt.Printf("Handling requests for user with ID %d\n", userId)
}
