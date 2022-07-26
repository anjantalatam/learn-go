package main

import "fmt"

func main() {

	// cards := newDeckFromFile("my_cards")

	// cards.print()
	// cards.shuffle()
	// cards.print()

	cards := newDeck()
	fmt.Println(len(cards), cards)
	a, b := cards.deal1(5)
	fmt.Println(len(cards), a, b, cards)
	fmt.Println(cards)
}
