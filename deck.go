package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create A Deck
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Create a file
func newDeckFromFile(name string) deck {
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Error:", err)
	}
	s := strings.Split(string(bytes), ",")
	return deck(s)
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Save cards to file on disk
func (d deck) saveToFile(name string) error {
	return os.WriteFile(name, []byte(d.toString()), 0666)
}

func (d deck) print() {
	for i, name := range d {
		fmt.Println(i, name)
	}
}
