package main

import (
	d "cards-project/m/v2/deck"
	"log"
)

func main() {
	deck := d.NewDeck()

	deck.Shuffle()

	hand, remainingCards := d.Deal(deck, 5)

	hand.Print()
	remainingCards.Print()

	err := deck.SaveToFile("my_deck")

	if err != nil {
		log.Panic("There was an error to save deck to the deck.txt file")
	}
}
