package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewCard(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected lenght of new deck = 52, found = %d", len(d))
	}

	//"Spades", "Hearts", "Diamonds", "Clubs"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck[0] = Ace of Spades, found = %s", d[0])
	}
	if d[13] != "Ace of Hearts" {
		t.Errorf("Expected deck[13] = Ace of Hearts, found = %s", d[13])
	}
	if d[26] != "Ace of Diamonds" {
		t.Errorf("Expected deck[26] = Ace of Diamonds, found = %s", d[26])
	}
	if d[39] != "Ace of Clubs" {
		t.Errorf("Expected deck[0] = Ace of Clubs, found = %s", d[39])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("./_deckTesting")

	d := newDeck()
	d.saveToFile("./_deckTesting")

	loadedDeck := newDeckFromFile("./_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Length of the deck loaded from file = %d, expected length = 52", len(loadedDeck))
	}

	os.Remove("./_deckTesting")
}

func ExampleTestPerm() {
	for _, value := range perm(4) {
		fmt.Println(value)
	}
	//Unordered Output: 0
	//1
	//2
	//3
}
