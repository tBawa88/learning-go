package main

import (
	"fmt"
	"time"
)

/*
	As soon as the 'done' channel closes, drain it
*/

func printProgress(fileSize <-chan int64, tick <-chan time.Time) {
	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)

		case <-done:
			for range done {
				//range on a channel actually executes a receive from the channel
				// so this empty for loop will drain the closes channel
			}
		}

	}
	fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)
}
