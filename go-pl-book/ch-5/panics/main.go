package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

// When a panic occurs, all defers calls that are pending are executed immidiately. (the defer stack is popped)
// and the stack trace is dumped to the stdout
