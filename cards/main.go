package main

func main() {
	cards, _ := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	//hand, cards := deal(cards, 5)

}

func newCard() string {
	return "Five of Diamonds"
}
