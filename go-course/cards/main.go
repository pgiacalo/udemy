package main

import "fmt"

func main() {

	cards := newDeck()
	cards.print()
	fmt.Println("----------------------------------")
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")

	// newCards := newDeckFromFile("my_cards")
	// newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
