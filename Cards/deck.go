package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
func newDeck() (cards deck) {
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return
}

// Shuffles the deck by swapping each element with another element of random index
func (d deck) shuffle() {
	sec := time.Now().UnixNano()
	source := rand.NewSource(sec)
	r := rand.New(source)

	for i := 0; i < len(d); i++ {
		r := r.Intn(len(d))
		d[i], d[r] = d[r], d[i]
	}
}

// Creates 2 "new" slices from  1 slice and returns them
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Takes a deck and saves to a file on HDD
func (d deck) saveToFile(filename string) error {
	b := []byte(d.toString())
	err := os.WriteFile(filename, b, 0666)
	return err
}

// Reads the file, splits the string into a [] string and returns a deck
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error :", err) //automatically calls os.Exit(1)
	}
	return strings.Split(string(data), ",")
}

// Converts []string to string using strings.Join() package method
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Just a dummy function created to write an Example test function
func perm(n int) (temp []int) {
	for i := 0; i < n; i++ {
		temp = append(temp, i)
	}
	return
}

// Constructs a single string from a deck ([] string), uses the named return to implicitly return the string
// func (d deck) toString() (s string) {
// 	cardSlice := []string(d)
// 	for i, value := range cardSlice {
// 		val := value + ","
// 		if i == len(d)-1 {
// 			val = value
// 		}
// 		s += val
// 	}
// 	return
// }
