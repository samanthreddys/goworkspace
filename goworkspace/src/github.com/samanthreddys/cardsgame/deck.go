package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Deck create a new type of Deck which is slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}

	}
	//fmt.Println(cards.toString())
	return cards
}

// Print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName) //bs is byteslice
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	//fmt.Println(s)
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
