package main

import "fmt"

// You can delay the execution of a function untill the surrounding functions returns

func printMessage(message string) {
	defer fmt.Println(message)
	fmt.Println("This line is written after the message")
}
