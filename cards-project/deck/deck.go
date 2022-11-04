package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var delimiter = ","

type Deck []string

func NewDeck() *Deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	deck := &Deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			*deck = append(*deck, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return deck
}

func NewDeckFromFile(fileName string) *Deck {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("There was an error when generating a deck from the file %s: %v", fileName, err)
		os.Exit(1)
	}

	deck := Deck(strings.Split(string(bytes), delimiter))
	return &deck
}

func (d *Deck) Print() {
	fmt.Printf("Printing a deck with %d cards\n\n", len(*d))

	for i, card := range *d {
		fmt.Printf("%d %s\n", i, card)
	}

	fmt.Println()
}

func Deal(d *Deck, length int) (*Deck, *Deck) {
	left := (*d)[:length]
	right := (*d)[length:]

	return &left, &right
}

func (d *Deck) Shuffle() {
	deck := *d
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range deck {
		newPosition := random.Intn(len(*d) - 1)
		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}

func (d *Deck) ToString() string {
	return strings.Join(*d, delimiter)
}

func (d *Deck) SaveToFile(fileName string) error {
	bytes := []byte(d.ToString())
	return os.WriteFile(fileName, bytes, 0666)
}
