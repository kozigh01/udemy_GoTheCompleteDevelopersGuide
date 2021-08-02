package deck_test

import (
	"os"
	"testing"

	"github.com/kozigh01/udemy_GoTheCompleteDevelopersGuide/cards/deck"
)

func TestNewDeckHas52Cards(t *testing.T) {
	const deckCount = 52
	deck := deck.NewDeck()

	if len(deck) != deckCount {
		t.Errorf("Expected deck length %v, but got %v", deckCount, len(deck))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	d := deck.NewDeck()

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be 'Ace of Spades', but was '%v'", d[0])
	}
}

func TestSaveToFileAndNewNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	defer func() { // cleanup after test
		os.Remove("_decktesting")
	}()

	deck1 := deck.NewDeck()
	deck1.SaveToFile("_decktesting")

	loadedDeck := deck.NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52, but got %v", len(loadedDeck))
	}
}
