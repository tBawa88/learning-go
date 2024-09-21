package main

import (
	"fmt"
	"time"
)

/*
	We want to loop over the filename slice, and process each file into a thumbnail.
	In this case, the task in each iteration is indipendent of the other, therefore we can delegate each iteration to a sperate goroutine
	To make this happen, we must initialize a buffered channel and make it's size equal to the slice length, so that each goroutine
	can send data to this channel as soon as it's done processing,  without getting blocked
*/

func main() {
	filenames := []string{"file9.png", "file6.png", "file5.png", "file4.png", "file3.png", "file2.png"}
	makeThumbnail(filenames)
}

func makeThumbnail(filenames []string) {
	ch := make(chan string, len(filenames))

	for _, f := range filenames {
		go func(file string) {
			processFile(f, ch)
		}(f)
	}

	// since we know the number of iterations, we will receive from the channel that exact number of times
	for range filenames {
		fmt.Println(<-ch)
	}
}

func processFile(file string, chh chan<- string) {
	time.Sleep(time.Second * 2)
	chh <- file + " processed "
}
