package main

import "fmt"

// Functions without a name
//Used for quick execution or to pass it as argument to another function

func printGreeting(greeting string) {
	func(greeting string) {
		fmt.Println("Hello there ", greeting)
	}(greeting)
}
