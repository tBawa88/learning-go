package main

import (
	"fmt"
	"time"
)

/*
	time.Tick(interval time.duration) returns a channel, on which it periodically sends events like a metronome
	The interval is defined by the duration passed as an argument
*/

func main() {
	fmt.Println("Commencing launch sequence")
	tick := time.Tick(time.Second * 1)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("LIFT OFF !!!!!!!")
}
