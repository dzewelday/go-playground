package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected: 16, Actual: %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: Ace of Spades, Actual: %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected: Four of Clubs, Actual: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const file = "_decktest"
	os.Remove(file)

	deck := newDeck()
	err := deck.saveToFile(file)
	if err != nil {
		t.Errorf("Failed to save deck to new file on disk, %v", err)
	}

	loadedDeck := newDeckFromFile(file)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected: 16, Actual: %v", len(loadedDeck))
	}
	if loadedDeck.toString() != deck.toString() {
		t.Error("Failed to create deck from file")
	}
	os.Remove(file)
}
