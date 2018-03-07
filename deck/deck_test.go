package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card of Two of Spaces, got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Expected last card of Ace of Spades, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestTestFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
