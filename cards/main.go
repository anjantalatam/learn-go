package main

import "fmt"

func main() {
	// var cards1 = []string{"hi","hello"}
	cards := []string{"Ace of diamonds", "King of Hearts"}
	cards = append(cards, "Queen of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
