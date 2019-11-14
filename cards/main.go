package main

func main() {
	/* cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Diamonds") */

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// remainingCards.print()
	// cards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := readAndGetNewDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// Type Conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}

/* func newCard() {
	return "Five of Diamonds"
}
*/
