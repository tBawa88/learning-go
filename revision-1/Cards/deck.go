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

var cardSuits = []string{"Hearts", "Spades", "Clubs", "Diamonds"}
var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "King", "Queen"}

func newDeck() (d deck) {
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return
}

func (d *deck) dealHand(size int) (deck, deck) {
	return (*d)[size:len(*d)], (*d)[0:size]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) saveToFile(filename string) {
	b := []byte(d.toString())
	os.WriteFile(filename, b, 0666)
}

func readFromFile(filename string) (hand deck) {
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading hand from file")
		log.Fatal(err)
	}

	return strings.Split(string(b), ",")
	// return byteToStringSlice(b)
}

func (d deck) shuffle() {
	for i := range d {
		randIdx := generateRandom(len(d))
		d[i], d[randIdx] = d[randIdx], d[i]
	}
}

// helper function to convert a deck ([] string) to a single string, using the strings package from go stl
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// creating a new source, then using that source to create a new Rand type object, then using that Rand to generate a random int
func generateRandom(limit int) int {
	source := rand.NewSource(time.Now().UnixMicro())
	randObj := rand.New(source)
	return randObj.Intn(limit)
}
