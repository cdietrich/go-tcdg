package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 32 {
		t.Errorf("Expected deck length of 32, but got %v", len(d))
	}
	if d[0] != "Seven of Clubs" {
		t.Errorf("Expected first card of Seven of Clubs, but got %v", d[0])
	}
	if d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Expected last card of Ace of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const f = "_testfile"
	r := func() {
		err := os.Remove(f)
		if err != nil {
			fmt.Println(err)
		}
	}
	r()
	defer r()
	d := newDeck()
	d.saveToFile(f)
	d2 := newDeckFromFile(f)
	if len(d) != len(d2) {
		t.Errorf("Expected %v values loaded, but got %v", len(d), len(d2))
	}
	for i, v := range d {
		if d2[i] != v {
			t.Errorf("Card %v does not match expected %v", d2[i], v)
		}
	}

}
