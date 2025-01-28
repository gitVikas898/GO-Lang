package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting)) // Type Conversion in Go

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards1 := newDeck()
	cards1.shuffle()
	cards1.print()

}
