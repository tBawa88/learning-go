package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	userId := 10

	// make an empty context
	// ctx := context.Background()

	// instead of an empty context, we can also make a context with value
	ctx := context.WithValue(context.Background(), userId, "bar")
	val, err := fetchUserId(ctx, userId)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("value = ", val)
	fmt.Printf("time taken = %s\n", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserId(ctx context.Context, userId int) (int, error) {
	value := ctx.Value(userId)
	fmt.Println("user value = ", value.(string))

	// creating a new context from the parent context passed to this function
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel() // defering cancel is done to prevent leaking of context resources

	// channel to communicate with the goroutine
	respch := make(chan Response)

	// starting a goroutine to offload this slow service
	go func(ch chan Response) {
		val, err := SlowThirdPartyService()
		ch <- Response{
			value: val,
			err:   err,
		}
	}(respch)

	// now we listen to 2 channels inside the for-select loop.
	// 1st channel is the the respch from which we expect the response to come out
	// 2nd channel is returned by this function call ctx.Done() to signify if the context had been closed or not

	// since we're using a context with a timeout, we can be guaranteed that the request to fetchUserId will not take longer than 300 ms
	for {
		select {
		case res := <-respch:
			return res.value, res.err
			// this is where the defered call to cancelFunc() comes in handy
			//since it is possible that we might return before the context's timeout has completed
			// defering it ensures that it will be closed in that case
		case <-ctx.Done():
			return 0, fmt.Errorf("third party response took too long")
		}
	}

}

// an outside 3rd party service which is unpredictable and could make performance of our web app slow
// we can use context to control this behaviour and make it "deterministic" instead of unpredictable
func SlowThirdPartyService() (int, error) {
	time.Sleep(time.Millisecond * 250)
	return 666, nil
}
