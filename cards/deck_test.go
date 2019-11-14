package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if 52 != len(d) {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if "Ace of Spades" != d[0] {
		t.Errorf("Expected first card as Ace of Spades, but got %v", d[0])
	}

	if "King of Clubs" != d[len(d)-1] {
		t.Errorf("Expected first card as King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndReadAndGetNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readAndGetNewDeckFromFile("_decktesting")

	if 52 != len(loadedDeck) {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_ducktesting")
}
