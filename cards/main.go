package main

import "fmt"

func main() {
// declare variable card
	// var card string = "Ace of Spades"
	// or
	// card := "Ace of Spades"
// reassign card value
	// card = "Five of Diamonds"
// function and return types
	// card := newCard()
// slices and loops
	cards := []string{newCard(), newCard()}

	for i, card := range cards {
		fmt.Println(i, card)
	}
}


func newCard() string {
	return "Five of Diamonds"
}
