package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		for i := range d {
			randIndex := rand.Intn(len(d))
			d[i], d[randIndex] = d[randIndex], d[i]
		}
	}
}

func (d Deck) SaveToFile(filename string) error {
	bytes := []byte(d.ToString())
	return ioutil.WriteFile(filename, bytes, 0666)
}

func NewDeckFromFile(filename string) Deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic("Mayday, mayday, mayday...")
		// return NewDeck()
	}

	cards := strings.Split(string(bytes), "|")
	return Deck(cards)
}

func (d Deck) DealHand(handSize int) (hand Deck, remainingDeck Deck) {
	rand.Seed(time.Now().UnixNano())
	hand = Deck{}
	remainingDeck = append(Deck{}, d[:]...)

	for i := 0; i < handSize; i++ {
		randomIndex := rand.Intn((len(remainingDeck) - 1) - i)
		hand = append(hand, remainingDeck[randomIndex])
		switch randomIndex {
		case 0:
			remainingDeck = remainingDeck[1:]
		case len(d):
			remainingDeck = remainingDeck[:(len(d) - 1)]
		default:
			remainingDeck = append(remainingDeck[:randomIndex], remainingDeck[(randomIndex+1):]...)
		}
	}

	remainingDeck = append(Deck{}, remainingDeck[:]...)
	return hand, remainingDeck
}

func (d Deck) NewCard() string {
	return "Five of Diamonds"
}

func (d Deck) Print() {
	fmt.Println("Cards in the deck:")
	for i, card := range d {
		fmt.Printf("   %02d, %v\n", i, card)
	}
}

func NewDeck() Deck {
	deck := Deck{}
	suits := [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := [13]string{
		"Ace", "King", "Queen", "Jack",
		"10", "9", "8", "7", "6", "5",
		"4", "3", "2",
	}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return deck
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), "|")
}
