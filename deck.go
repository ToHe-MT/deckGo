package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new etype of 'deck'
// Which is a slice of strings

type deck []string

// NEW DECK
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNumbers := []string{"Ace",
		// "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardNumbers {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newI := r.Intn(len(d) - 1)
		d[i], d[newI] = d[newI], d[i]
	}
}
