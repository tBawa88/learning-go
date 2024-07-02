package main

import "fmt"

// import "fmt"

func main() {

	//cards is now a of type deck (deck is a type which extends []string 'a list of string')
	cards := deck{"Jack of Spades", newCard(), "Queen of Hearts"}
	cards = append(cards, "8 of Clubs")

	cards.print()
	cards.addCard(newCard())
	cards.addNewCard(newCard())
	fmt.Println("Added a new card----")
	cards.print()

}

// Returns a new card
func newCard() string {
	return "Five of Diamonds"
}
