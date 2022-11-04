package deck

import (
	"os"
	"testing"
)

func TestCardEqual(t *testing.T) {
	c := &Card{
		suit:  "Spades",
		value: "Ace",
	}

	c1 := &Card{
		suit:  "Spades",
		value: "Ace",
	}

	c2 := &Card{
		suit:  "Clubs",
		value: "Ace",
	}

	c3 := &Card{
		suit:  "Hearts",
		value: "Ace",
	}

	if !c.Equal(*c1) {
		t.Errorf("Expected %v to equal %v, but its not equal", c, c1)
	}

	if c2.Equal(*c3) {
		t.Errorf("Expected %v to not equal %v, but its equal", c, c1)
	}
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	first := (*deck)[0]
	last := (*deck)[len(*deck)-1]

	expectedFirst := &Card{
		suit:  "Spades",
		value: "Ace",
	}
	expectedLast := &Card{
		suit:  "Clubs",
		value: "King",
	}

	if len(*deck) != 52 {
		t.Errorf("Expected deck of 52, but got %d", len(*deck))
	}

	if !first.Equal(*expectedFirst) {
		t.Errorf("Expected first card of %s, but got %s", expectedFirst, first)
	}

	if !last.Equal(*expectedLast) {
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
