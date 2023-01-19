package main

import (
	"os"
	"testing"
)

func TestSaveToFileAndNewFile(t *testing.T) {
	os.Remove("_testing.go")
	cards := newDeck()
	cards.saveToFile("decktesting")
	d := newDeckFromFile("decktesting")
	if len(d) != 16 {
		t.Errorf("Expected length of deck %v", len(d))
	}
	os.Remove("_testing.go")
}
