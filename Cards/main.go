package main

import "fmt"

func main() {

	cardDeck := newDeck()
	// fmt.Println("Current Deck ===")
	// cardDeck.print()

	hand, cardDeck := cardDeck.deal(10) //redeclaring cardDeck is allowed as long as we're declaring one more variable for the fist time along with it

	// fmt.Println("==== New hand ======")
	// hand.print()
	// fmt.Println("----- New Deck -----")
	// cardDeck.print()

	// s := hand.toString()
	s := hand.ToString()
	fmt.Println(s)

	err := hand.saveToFile("./cardsInfo")
	if err != nil {
		fmt.Println("Error saving to file to HD", err)
	} else {
		fmt.Println("Card info saved to file!!")
	}

	cards := newDeckFromFile("./cardsInfo")
	cards.print()
}
