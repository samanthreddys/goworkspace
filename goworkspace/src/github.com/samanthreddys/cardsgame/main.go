package main

func main() {
	cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//	fmt.Println("break")
	//remainingDeck.print()

	//fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	cards.saveToFile("DeckofCards")
	newDeckromFile("DeckofCards")
	//r.print()

}
