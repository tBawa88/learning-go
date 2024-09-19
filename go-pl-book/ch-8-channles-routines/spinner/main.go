package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) // this function will be ran by a separate goroutine while the main goroutine is busy executing this

	fiB := fib(45) // slow function call

	fmt.Printf("fibonacci = %d\n", fiB)
	// as soon as the main goroutine exits the program, all goroutines are terminated

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // '\r' is called carriage return moves the cursor back to the beginning of the line
			time.Sleep(delay)     // this is a blocking call, when the routine hits this call, it has to sleep for delay duration
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
