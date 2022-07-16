package man

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)

}

func newCard() string {
	return "Kind of Diamonds"
}
