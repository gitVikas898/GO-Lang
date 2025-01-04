package main

import (
	"fmt"
)

func main() {

	cards := []string{newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades")

	nums := []int{10, 11, 12, 13, 14}

	for i, num := range nums {
		fmt.Println(i, num)
	}

	fmt.Println(cards)

	result := findSum()

	fmt.Println("The Sum is : ", result)

}

func newCard() string {
	return "Five of Diamonds"
}

func number() int {
	return 12
}

func findSum() int {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for num := range nums {
		sum = sum + num
	}

	return sum
}

// Array(fixed) & Slice(dynamic);

/*
	Every element in a slice must be aof same type
*/
