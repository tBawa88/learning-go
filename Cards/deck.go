package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type deck []string

var cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

// Prints the deck that calls this method
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Returns a new deck, which contains all 52 cards
func newDeck() deck {
	var freshDeck deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			freshDeck = append(freshDeck, card)
		}
	}
	return freshDeck
}

// Returns 2 decks, a hand and new deck (deck, deck)
func (d deck) deal(handSize int) (deck, deck) {
	//We're taking out cards from the existing deck and creating a hand. Therfore those cards must be removed from the exisiting deck
	newHand := d[:handSize]
	newDeck := d[handSize:]

	return newHand, newDeck
}

// Takes a deck and saves to a file on HDD
func (d deck) saveToFile(filename string) error {
	b := []byte(d.ToString())
	err := os.WriteFile(filename, b, 0666)
	return err
}

// Reads the file, splits the string into a [] string and returns a deck
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		//Option 1 - log the error, and return a brand newDeck() so as to satisfy the return type
		//Option 2 - Log the error, and exit the program entirely and prevent any further execution
		// fmt.Println("Error : ", err)
		// os.Exit(1) //either use this or use log.Fatal(err) -> it logs the error and automatically calls os.Exit()
		log.Fatal("Error :", err)
	}

	//data ([] byte) gets converted into string, then that string is getting split into a [] string
	//We can directly return a []string because deck is extending [] string so compiler can easily convert it into deck
	return strings.Split(string(data), ",")

}

// Constructs a single string from a deck ([] string), uses the named return to implicitly return the string
func (d deck) toString() (s string) {
	cardSlice := []string(d) //convert the deck into a string slice
	for i, value := range cardSlice {
		val := value + "," //add comma after each value so that when we deconstruct this string we can easily obtain the original slice
		if i == len(d)-1 { //if it's the last value, don't add the comma
			val = value
		}
		s += val
	}
	return
}

// Same funciton, but uses the inbuilt strings.Join() function from "strings"
func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}
