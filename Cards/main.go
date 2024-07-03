package main

import "fmt"

func main() {

	cards := newDeck()
	cards.shuffle()
	hand, _ := cards.deal(5)

	fmt.Println("Hand from a shuffled deck ------------")
	hand.print()
}
