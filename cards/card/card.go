package card

import (
	"fmt"
)

func NewCard() string {
	return "Five of Diamonds"
}

func PrintCards(deck []string) {
	fmt.Println("Cards in the deck:")
	for i, card := range deck {
		fmt.Printf("   %v, %v\n", i, card)
	}
}
