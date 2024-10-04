package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, Cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer Cancel() // since it will automatically be closed, we only need to call it if we want to end it early

	// start a goroutine and bind this context with it
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context timed out, Closing operation")
				wg.Done()
				return
			default:
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Operation running")
			}
		}
	}(ctx)

	wg.Wait()
}
