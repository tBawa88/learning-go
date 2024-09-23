package main

import (
	"fmt"
	"time"
)

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
		}

	}
	fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)
}
