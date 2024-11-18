package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// create a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%s of %s", value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

// get all deck
func (d deck) print() {
	for index, card := range d {
		fmt.Println(
			index,
			card,
		)
	}
}

// post a hand of card
func deal(d deck, handsize int) (deck, deck) {
	// slice default behavior - slice[i=0:i-1] ~ deck[:3]-> 0,1,2
	// slice default behavior - slice[i=0:i-1] ~ deck[3:]-> 3, ...end
	return d[:handsize], d[handsize:]
}

// post a hand of card
func (d deck) mDeal(handsize int) (deck, deck) {
	// slice default behavior - slice[i=0:i-1] ~ deck[:3]-> 0,1,2
	// slice default behavior - slice[i=0:i-1] ~ deck[3:]-> 3, ...end
	return d[:handsize], d[handsize:]
}

// turn a deck to string
func (d deck) to_string() string {
	return strings.Join([]string(d), ",")
}

// save file
func (d deck) saveToFile(filename string) error {
	//0666 octave number to give file permissions
	return os.WriteFile(filename, []byte(d.to_string()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		new_position := r.Intn(len(d) - 1)
		d[i], d[new_position] = d[new_position], d[i]
	}
}
