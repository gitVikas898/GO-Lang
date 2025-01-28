package main

import (
	"fmt"
	"os"
	"strings"
)

//create a new type of deck
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuit := []string{"Spade", "Heart", "Club", "Diamond"}

	cardsValue := []string{"Ace", "2", "3", "4"}

	for _, suit := range cardsSuit {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// utility function to print
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// deal function
func deal(pack deck, handSize int) (deck, deck) {
	return pack[:handSize], pack[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
