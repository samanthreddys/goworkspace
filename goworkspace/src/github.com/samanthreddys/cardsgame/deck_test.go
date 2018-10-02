package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		// TODO: Add test cases.
		//{"SC1", deck{"Queen of Hearts", "Six of Diamonds", "Nine of Spades", "Nine of Diamonds", "Ten of Spades", "Two of Clubs", "King of Diamonds", "Eight of Hearts", "Ace of Spades", "Jack of Clubs", "Three of Hearts", "King of Spades", "Nine of Hearts", "Five of Diamonds", "Ten of Diamonds", "Four of Diamonds", "Four of Spades", "Four of Hearts", "Eight of Diamonds", "King of Hearts", "Two of Diamonds", "Seven of Clubs", "Seven of Hearts", "Three of Spades", "Queen of Spades", "Five of Hearts", "Queen of Clubs", "Ace of Hearts", "Seven of Diamonds", "Ace of Clubs", "Two of Hearts", "Eight of Spades", "King of Clubs", "Jack of Spades", "Six of Spades", "Five of Spades", "Three of Clubs", "Five of Clubs", "Three of Diamonds", "Nine of Clubs", "Ten of Hearts", "Queen of Diamonds", "Six of Hearts", "Ace of Diamonds", "Six of Clubs", "Four of Clubs", "Eight of Clubs", "Ten of Clubs", "Seven of Spades", "Two of Spades", "Jack of Diamonds", "Jack of Hearts"}},
		{"Sc2", deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_newDeckLength(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Length is 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])

	}
}

func Test_deck_print(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
		{"Sc2", deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.print()
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck but got %v", len(loadedDeck))
	}
	/* type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"decktest", deck{"Hello"}, "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("deck.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	} */
}

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffle()
		})
	}
}
