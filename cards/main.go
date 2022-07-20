package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	newDeckFromFile("my_card")

}
