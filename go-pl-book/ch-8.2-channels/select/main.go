package main

import "fmt"

/*
	A select statement, kinda like switch case statemnt defines multiple cases. Each case represents a communication between a routine and a channel
	(A send or a receive)

	select waits untill communication for one of the cases is ready to proceed, and executes it's code.
	The first case, which is ready to communicate is executed by the select statement

	IF, MORE THAN ONE CASES ARE READY TO COMMUNICATE, select picks one at RANDOM. To ensure that each case has equal chance to communicate to the channel
*/

func main() {

	// A buffered channel which can hold 1 element at time
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {

		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			//value sent into the channel
		}
	}

	fmt.Println("Now printing the second select which is random")
	// A channel with bigger buffer to display *non-deterministic* execution when more than one channels are ready to communicate
	ch2 := make(chan int, 2)
	for i := 0; i < 10; i++ {

		select {
		case x := <-ch2:
			fmt.Println(x)
		case ch2 <- i:
			//value sent into the channel
		}
	}
}

/*
	EXPLANATION
	1.On the first iteration, the channel is empty, therefore there cannot be any receive, hence case x := <- chan: is not ready to communicate
	2. BUT on the fist iteration, since the channel is empty, we can send value to the channel, therefore 0 is sent to the channel

	3. One the next iteration, the channel is full, therefore the 2nd case cannot communicate, but the first case can
	4. This way the loop prints out even values, taken from the alternate iteration
*/
