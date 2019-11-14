package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	/* 	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	} */

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	fmt.Println("Shuffling cards")

	// for index := range d {
	// 	randomPosition := rand.Intn(len(d) - 1)
	// 	d[index], d[randomPosition] = d[randomPosition], d[index]
	// }

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		randomPosition := r.Intn(len(d) - 1)
		d[index], d[randomPosition] = d[randomPosition], d[index]
	}
}

func readAndGetNewDeckFromFile(fileName string) deck {
	byteStream, err := ioutil.ReadFile(fileName)
	if nil != err {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteStream), ","))
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.toByteSlice(), 0666)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
