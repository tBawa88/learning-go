package main

import "fmt"

// UNI-DIRECTIONAL channels
// Usually different go routines only need to access only one side of the channel
// they either need to receive the value coming out the channel, or send a value to the channel
// therefore Go provides us with uni directional channels

// a goroutine that receives a sending type channel, cannot receive value from it (results in compile time error)
// a goroutine that reives a receiving type channel, cannot send value to it
// close() can only be called by the goroutine that has received sending type channel (since 'close' is meant to signify no more sending )

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals, 10)
	go squarer(squares, naturals)

	printer(squares)

}

// this is how we define a seding type channel 	chan <- int (int value going into the channel)
func counter(output chan<- int, n int) {
	for n > 0 {
		output <- n
		n--
	}
	close(output)
}

// the input chan is a receiver type channle (squares cannot send any value to this channel)
func squarer(output chan<- int, input <-chan int) {
	for nat := range input {
		output <- nat * nat
	}
	close(output)
}

// ran by main goroutine and accepts a receiver type channel
func printer(input <-chan int) {
	for n := range input {
		fmt.Println(n)
	}
}
