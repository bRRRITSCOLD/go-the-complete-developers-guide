package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()

	cards.shuffle()

	dealtCards, remainingCards := deal(cards, 3)

	dealtCards.savetoFile("dealt_cards")

	remainingCards.savetoFile("remaining_cards")

	readDealtCards, err := newDeckFromFile("dealt_cards")
	if err != nil {
		os.Exit(1)
	}

	readRemainingCards, err := newDeckFromFile("remaining_cards")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("------ DEALT CARDS ------")
	readDealtCards.print()

	fmt.Println("------ REMAINING CARDS ------")
	readRemainingCards.print()
}
