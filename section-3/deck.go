package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{
		"Spades",
		"Hearts",
		"Diamonds",
		"Clubs",
	}

	cardValues := []string{
		"Queen",
		"King",
		"Jack",
		"Ten",
		"Nine",
		"Eight",
		"Seven",
		"Six",
		"Five",
		"Four",
		"Three",
		"Two",
		"Ace",
	}

	d := deck{}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			d = append(d, cardValue+" of "+cardSuit)
		}
	}

	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, numberOfCards int) (dealtCards deck, newDeck deck) {
	return d[:numberOfCards], d[numberOfCards:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return deck{}, err
	}
	s := strings.Split(string(bs), ",")
	return deck(s), nil
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
