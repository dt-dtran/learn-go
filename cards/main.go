package main

import "fmt"

func main() {
// declare variable card
	// var card string = "Ace of Spades"
	// or
	// card := "Ace of Spades"
// reassign card value
	// card = "Five of Diamonds"

	card := newCard()

	fmt.Println(card)
}


func newCard() string {
	return "Five of Diamonds"
}
