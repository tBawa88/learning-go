package main

import "fmt"

// import "fmt"

func main() {
	// d := newDeck()
	// d, hand := d.dealHand(13)

	// d.print()
	// fmt.Println("Printing the new hand")
	// hand.print()

	// //let's save the hand to the file
	// hand.saveToFile()

	// //fetching the hand from file

	// newHand := readFromFile("./deckInfo")
	// fmt.Println("Fetching hand from file ........")
	// newHand.print()

	d := newDeck()
	d.print()

	fmt.Println("Shuffling the deck ....")
	d.shuffle()
	d.print()
}
