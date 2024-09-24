package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

/*
	This is the version after directory_trav_3
	In this we're attempting to cancel all the running gorutines if the user presses enter in ther terminal stdin
	Closing the goroutines(when they're not needed) is just as important as starting them
*/

/*
	Approach:
	- We create a new channel called 'done'. IT will never receive any values, but closing this will signify the end of all running goroutines
	- Create a helper function called 'cancelled' which checks if the current channel is closes
	- Inside the printProgess, add another case in the select statement which listens to closing of this channel
		- This case will drain the channel as soon as it closes, (so that susbsequent goroutines can start getting zero values from it's reads)
	- Inside the walkDir() function, we check the cancellation and return the function call, ending the goroutine
	- It's always better to check for cancellation in more than one places, especially a few important places.
	- One such place in this program is the acquisition of the semaphore, if we check the cancellation before that, we can stop the goroutine from opening the directory
		- could end up imporving cancel times
*/

var vFlag = flag.Bool("v", false, "verbose progress of directory traversal")
var dirName = flag.String("dir", ".", "name of the directory to be traversed")
var done = make(chan struct{})

// helper function
func cancelled() bool {
	/*
		We cannot do something like this
		_, ok := <- done
		return ok		// since it will be a blocking call (ain't none sending to this channel)
	*/
	select {
	case <-done: // once the channel is closed, all receives proceed immidiatley, meaning this case will only run if the channel is closed in our case
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()
	fileSize := make(chan int64)
	wg := &sync.WaitGroup{}

	// starting the first goroutine
	wg.Add(1)
	go func() {
		walkDir(*dirName, fileSize, wg)
		close(fileSize)
	}()

	// closer goroutine
	go func() {
		wg.Wait()
		close(fileSize)
	}()

	// cancler goroutine
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var tick <-chan time.Time
	if *vFlag {
		fmt.Println("verbose flag passed ")
		tick = time.Tick(time.Millisecond * 100)
	}
	printProgress(fileSize, tick)
}
