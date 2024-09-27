package main

import (
	"fmt"
	"sync"
)

/*
	A data race occurs when multiple goroutines access a shared variable and one of them accesses is a write operation
	In these cases, it's almost impossible to predict the final state of the shared variable

	In the main function below, we're starting 99999 goroutines almos at once as we go through the for loop
	Now it's not guaranteed that all go routines will get the latest value of the counter. Since the sample case is so large, most of 'em will get an older value of counter
	while some other goroutine has already updated the real value, making current goroutine's write opration being lost.
	It's impossible to predict the order of execution of goroutines
*/

func main() {
	wg := sync.WaitGroup{}

	counter := 0
	for i := 0; i < 99999; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Final value of counter  = ", counter)
}
