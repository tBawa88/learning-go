package main

import "fmt"

// Create a new 'type' of deck
// which is nothing but a slice of strings
type deck []string

// A function that prints all the cards present inside the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// // Changes my by receiver that receives by value do not reflect on the original variable
// func (d deck) addCard(card string) {
// 	d = append(d, card)
// }

// This is called pointer receiver. Any object that calls this method will have it's address passed to this and stored inside the variable d
func (d *deck) addNewCard(card string) {
	*d = append(*d, card)
}
