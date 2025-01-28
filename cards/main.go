package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting)) // Type Conversion in Go

	cards := newDeck()
	cards.saveToFile("my_cards")

}
