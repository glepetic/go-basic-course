package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	// Assemble
	expectedLen := 52
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "King of Clubs"

	// Act
	target := newDeck()

	// Assert
	targetLen := len(target)
	if targetLen != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, targetLen)
	}

	targetFirstCard := target[0]
	if targetFirstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, targetFirstCard)
	}

	targetLastCard := target[targetLen-1]
	if targetLastCard != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, targetLastCard)
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	// Assemble
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := deck{"Ace of Hearts", "Jack of Clubs", "Queen of Diamonds", "King of Spades"}
	expectedLen := len(deck)
	expectedFirstCard := deck[0]
	expectedLastCard := deck[expectedLen-1]

	// Act
	deck.saveToFile(fileName)
	target := newDeckFromFile(fileName)

	// Assert
	targetLen := len(target)
	if targetLen != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, targetLen)
	}

	targetFirstCard := target[0]
	if targetFirstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, targetFirstCard)
	}

	targetLastCard := target[targetLen-1]
	if targetLastCard != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, targetLastCard)
	}

}
