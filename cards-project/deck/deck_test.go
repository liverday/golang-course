package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	first := (*deck)[0]
	last := (*deck)[len(*deck)-1]

	expectedFirst := "Ace of Spades"
	expectedLast := "King of Clubs"

	if len(*deck) != 52 {
		t.Errorf("Expected deck of 52, but got %d", len(*deck))
	}

	if first != expectedFirst {
		t.Errorf("Expected first card of %s, but got %s", expectedFirst, first)
	}

	if last != expectedLast {
		t.Errorf("Expected last card of %s, but got %s", expectedLast, last)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktest"
	deck := NewDeck()
	deck.SaveToFile(fileName)
	defer os.Remove(fileName)

	deckFromFile := NewDeckFromFile(fileName)

	if len(*deckFromFile) != 52 {
		t.Errorf("Expected deck of 52, but got %d", len(*deck))
	}
}
