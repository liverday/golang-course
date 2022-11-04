package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var delimiter = ","

type Card struct {
	suit  string
	value string
}

func (c *Card) Equal(other Card) bool {
	return c.suit == other.suit && c.value == other.value
}

type Deck []*Card

func NewDeck() *Deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	deck := &Deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			*deck = append(*deck, &Card{
				suit:  suit,
				value: value,
			})
		}
	}

	return deck
}

func getSuitAndValueFromPlainCard(card string) (string, string) {
	arr := strings.Split(card, " of ")

	return arr[0], arr[1]
}

func NewDeckFromFile(fileName string) *Deck {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("There was an error when generating a deck from the file %s: %v", fileName, err)
		os.Exit(1)
	}

	cardValues := strings.Split(string(bytes), delimiter)

	deck := &Deck{}

	for _, card := range cardValues {
		suit, value := getSuitAndValueFromPlainCard(card)

		*deck = append(*deck, &Card{
			suit:  suit,
			value: value,
		})
	}

	return deck
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

func (d *Deck) ToArrayOfStrings() []string {
	arr := make([]string, len(*d))

	for i, card := range *d {
		arr[i] = fmt.Sprintf("%s of %s", card.value, card.suit)
	}

	fmt.Printf("Arr %v\n", arr)

	return arr
}

func (d *Deck) ToString() string {
	return strings.Join(d.ToArrayOfStrings(), delimiter)
}

func (d *Deck) SaveToFile(fileName string) error {
	bytes := []byte(d.ToString())
	return os.WriteFile(fileName, bytes, 0666)
}
