package main

func main() {
// 1. declare variable card
	// var card string = "Ace of Spades"
	// or
	// card := "Ace of Spades"
// 2. reassign card value
	// card = "Five of Diamonds"
// 3. function and return types
	// card := newCard()
// 4. slices and loops
	// cards := []string{newCard(), newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
// 5. type method: print()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
// 6. type methid: newDeck()
	cards := newDeck()
	cards.print()
}

// 3. function and return types
// func newCard() string {
// 	return "Five of Diamonds"
// }
