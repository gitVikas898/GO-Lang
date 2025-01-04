package main

import "fmt"

//create a new type of deck
//which is a slice of strings

type deck []string

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

type nums []int

func (n nums) print() {
	for num := range n {
		fmt.Println(n[num])
	}
}
