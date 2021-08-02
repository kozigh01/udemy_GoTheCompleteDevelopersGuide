package deck_test

import (
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
