package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Context and Cancelfunc
	ctx, Cancel := context.WithCancel(context.Background())

	// start a goroutine that listens for context cancellation
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("Context has been cancelled. It's career is done")
	}(ctx)

	// simulate a delay
	time.Sleep(time.Second * 2)
	Cancel()

	// holding maingoroutine so that other goroutine can finish
	time.Sleep(time.Millisecond * 300)
}

/*
	So imagine, we can start a goroutine which handles some operation
	We can create special context for that operation and bind the goroutine with that context
	Now we have the power to terminate that operation from outside the goroutine
*/
