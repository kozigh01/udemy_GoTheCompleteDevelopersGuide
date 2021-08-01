package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []string

func (d Deck) DealHand(handSize int) (hand Deck, deck Deck) {
	rand.Seed(time.Now().UnixNano())
	hand = Deck{}

	for i := 0; i < handSize; i++ {
		randomIndex := rand.Intn((len(d) - 1) - i)
		hand = append(hand, d[randomIndex])
		switch randomIndex {
		case 0:
			d = d[1:]
		case len(d):
			d = d[:(len(d)-1)]
		default:
			d = append(d[:randomIndex], d[(randomIndex+1):]...)
		}
	}

	return hand, append(Deck{}, d[:len(d)]...)
}

func (d Deck) NewCard() string {
	return "Five of Diamonds"
}

func (d Deck) Print() {
	fmt.Println("Cards in the deck:")
	for i, card := range d {
		fmt.Printf("   %v, %v\n", i, card)
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
			deck = append(deck, fmt.Sprintf("%s %s", suit, value))
		}
	}

	return deck
}