package main

import (
	"os"
	"testing"
)

var size = 52
var firstCard = "Ace of Hearts"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != size {
		t.Errorf("Expected new deck size to be 52, got %v", len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected first card to be Ace of Hearts, got %s", d[0])
	}

}

func TestFileIO(t *testing.T) {
	fileName := "_deckTesting"
	os.Remove(fileName)
	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := readFromFile(fileName)

	if len(loadedDeck) != size {
		t.Errorf("Expected new deck size to be 52, got %v", len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected first card to be Ace of Hearts, got %s", d[0])
	}

	os.Remove(fileName)

}
