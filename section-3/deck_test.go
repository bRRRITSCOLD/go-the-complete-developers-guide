package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Queen of Spades" {
		t.Errorf("Expected first card to be Queen of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card to be Ace of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_deck-test-save-to-file"

	err := os.Remove(filename)

	d := deck{
		"hello",
		"world",
	}

	d.savetoFile(filename)

	dir, err := ioutil.ReadDir("./")
	if err != nil {
		t.Errorf("Failed to read dir for test purposes, received error %v", err.Error())
	}

	filenameFound := false
	for _, file := range dir {
		if file.Name() == filename {
			filenameFound = true
		}
	}
	if filenameFound == false {
		t.Errorf("Expected to find file %v after saveToFile, but found no file with said name", filename)
	}

	rd, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Failed to create new dec from file, received error %v", err.Error())
	}
	if len(rd) != 2 {
		t.Errorf("Expected number of cars in new deck from file to be 2, but go %v", len(rd))
	}

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to clean up %v, received error %v", filename, err.Error())
	}
}
