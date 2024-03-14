package main

import "fmt"

// create a new type of deck as a slice of string
type deck []string

// receiver is not needed
func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
