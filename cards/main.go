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
// 6. type method: newDeck()
	// cards := newDeck()
	// cards.print()
// 7. type method: deal()
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("hand")
	// hand.print()
	// fmt.Println("RE")
	// remainingCards.print()
// 8. type method: convert deck to []string
	// fmt.Println(cards.toString())
// 9. type method: save to file
	// cards.saveToFile("my_cards")
// 10. type method: get cards from file
	cards := newDeckFromFile("my_cards")
	cards.print()
}

// 3. function and return types
// func newCard() string {
// 	return "Five of Diamonds"
// }
