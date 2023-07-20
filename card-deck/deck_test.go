package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetNewCardDeck(t *testing.T) {
	d := newCardDeck()
	if len(d) != 28 {
		t.Errorf("Expected length of deck to be 28 but got %v", len(d))
	}
	if strings.Compare(d[0], "Six Of Spades") != 0 {
		t.Errorf("Expected first card to be 'Six of Spades' but got %v", d[0])
	}

	if strings.Compare(d[len(d)-1], "Seven Of Hearts") != 0 {
		t.Errorf("Expected last card to be 'Hearts Of Seven' but got %v", d[len(d)-1])
	}
}

func TestWriteToFile(t *testing.T) {
	d := newCardDeck()
	_, err := d.writeToFile("_testing")
	if err != nil {
		t.Errorf("Error in writing deck to file : %v", err)
	}
	os.Remove("_testing")
}
func TestGetDeckFromFile(t *testing.T) {
	d := newCardDeck()
	_, err := d.writeToFile("_testing")
	deckFromFile, err := getDeckFromFile("_testing")
	if err != nil {
		t.Errorf("Error getting file '%v' from hard drive!", "_testing")
	} else {
		for i := 0; i < len(deckFromFile); i++ {
			if strings.Compare(d[i], deckFromFile[i]) != 0 {
				t.Errorf("Data read from file is not same as written data")
				break
			}
		}
	}
	os.Remove("_testing")
}
