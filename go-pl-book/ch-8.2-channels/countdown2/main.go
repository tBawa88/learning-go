package main

import (
	"fmt"
	"os"
	"time"
)

/*
	MULTIPLEXING WITH A 'SELECT' STATEMENT
	Let's add the ability to abort the launch sequence if an enter key is pressed in the stdin
*/

func launch() {
	fmt.Println("LIFT OFF !!!!!!!")
}

func main() {

	abort := make(chan struct{})
	go func() {
		// os.Stdin satisfies Reader, therefore we can just call the Read function on it, it waits untill newline is pressed and it successfuly reads some data from stdin
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	/*
		Since we cannot listen to one channel first (because it's a blocking call untill the channel sends data)
		we need to listen to both of these channels at the same time,
		if a value comes out of abort, we need to exit the main goroutine
		This is done using a **select** statement
		Which looks a lot similar to a switch-case statement
	*/

	fmt.Println("Commencing launch sequence, Press return to abort")

	select {
	case <-time.After(time.Second * 10):
		launch()
	case <-abort:
		fmt.Print("Launch aborted")
		return
	}
	/*
		time.After() too returns a channel, to which it sends current timestamp, after the specified duration has passed
		A Select statement **waits** untill one of it's cases is ready communicate.
		It then executes the communication for that case and other communication never happen
		A select{} waits forever
	*/
}
