package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("LIFT OFF !!!!!!!")
}

// Making the program print the sequence and have the ability to shutdown
func main() {

	abort := make(chan int)
	tick := time.Tick(time.Second * 1)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 1
	}()

	fmt.Println("Commencing launch sequence, Press return to abort")
	for count := 10; count > 0; count-- {
		select {
		case <-tick:
			fmt.Println(count)

		case <-abort:
			fmt.Println("aborting launch")
			return

		}
	}
	launch()
}
