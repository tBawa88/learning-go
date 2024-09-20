package main

import "fmt"

// 3 different go routines communicating sending inputs to each other through a channel using it as a pipeline
// to CLOSE a channel, we make the sender goroutine call the close(chan) function which sets a flag that tells the receiver that no more value will come out of this channel
// but a closed channel can continue to supply 'zero values' if the receiver tries to access it
// therefore there are 2 ways to check if the channel has been closed or not
// value , ok := <- channel  	// the ok is a boolean which tells that this channel has been closed. Therefore the accessing routine can only get values untill the
// channel is drained (meaning all remaining values have been pulled out if the channel is a buffered type)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			n, ok := <-naturals // checking if the natural channel has been closed and "drained"(all input sent values have been received)
			if !ok {
				break
			}
			squares <- n * n
		}
		close(squares)
	}()

	// printer in main goroutine. When the channel is closed by a routine, the loop using 'range' to iterate over a channel also terminates
	// we don't need to obtain a boolean and check for it every time if the channel has been closed or not
	for sq := range squares {
		fmt.Println(sq)
	}

}
