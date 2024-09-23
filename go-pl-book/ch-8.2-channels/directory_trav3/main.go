package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

/*
	Directory traversal over large set of files is still too slow
	We can make this concurrent by starting a new goroutine for every new walkDir() call
	And use WaitGroup() to keep track of last goroutine execution

	SEMAPHORE : since this program creates thousands of goroutines at it's peak, we need to implement a semaphore
	to prevent it from opening too many files at once and running out of file descriptors

	Go look at just above dirents function
*/

var vFlag = flag.Bool("v", false, "verbose progress of directory traversal")
var dirName = flag.String("dir", ".", "name of the directory to be traversed")

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

	// starting the closer goroutine
	go func() {
		wg.Wait()
		close(fileSize)
	}()

	var tick <-chan time.Time
	if *vFlag {
		fmt.Println("verbose flag passed ")
		tick = time.Tick(time.Millisecond * 100)
	}
	printProgress(fileSize, tick)
}
