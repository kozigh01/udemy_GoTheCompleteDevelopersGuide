package main

import (
	"fmt"	

	"github.com/kozigh01/udemy_GoTheCompleteDevelopersGuide/cards/card"
)

func main() {
	fmt.Println("Cards Project")

	card1 := card.NewCard()
	fmt.Printf("Card: %q, type: %T\n", card1, card1)

	var cards [52]string
	cards[0] = "Ace of Hearts"
	cards[51] = "Two of Spades"
	// fmt.Printf("cards: %q (type: %T)\n", cards, cards)
	fmt.Println("Cards in the deck:")
	for i, card := range cards {
		fmt.Printf("   %v, %v\n", i, card)
	}
	
	// cards2 := []string{card.NewCard(), "King of Hearts"} // literal slice initialization
	cards2 := make([]string, 0, 52)  // use make if length and capacity are known
	// cards2[0] = "Ace of Hearts"  // throws error
	cards2 = append(cards2, "Ace of Hearts")
	cards2 = append(cards2, "King of Hearts")
	cards2[0] = "diff card"
	// fmt.Printf("cards2: %+q (type: %T, length: %v, capacity: %v)\n", cards2, cards2, len(cards2), cap(cards2))
	card.PrintCards(cards2)

	cards3 := make([]string, 52)
	cards3[0] = "Two of Spades"
	cards3[51] = "Ace of Clubs"
	// fmt.Printf("cards3: %+q (type: %T, length: %v, capacity: %v)\n", cards3, cards3, len(cards3), cap(cards3))
	// nicePrint(fmt.Sprintf("cards3: %+q (type: %T, length: %v, capacity: %v)\n", cards3, cards3, len(cards3), cap(cards3)))
	card.PrintCards(cards3)
}

func nicePrint(s string) {
	fmt.Println("     ----------")
	defer fmt.Println("    ----------")
	fmt.Print(s)
}