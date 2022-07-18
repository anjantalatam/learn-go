package main

func main() {
	// var cards1 = []string{"hi","hello"}
	cards := deck{"Ace of diamonds", "King of Hearts"}
	cards = append(cards, "Queen of Diamonds")

	cards.print()
}
