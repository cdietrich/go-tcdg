package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.saveToFile("demo.deck")
	cards = newDeckFromFile("demo.deck")
	cards.shuffe()
	cards.print()
	fmt.Println("++++++++++++++++")
	var hand deck
	noh := len(cards) / 4
	for i := 0; i < noh; i++ {
		hand, cards = deal(cards, 4)
		fmt.Println(hand)
	}
	dx := deck{"A", "B", "C"}
	dx.print()
	dx.saveToFile2("dx.txt")
	dx = newDeckFromFile2("dx.txt")
	dx.print()
}
