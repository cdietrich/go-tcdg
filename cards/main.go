package main

import (
	"fmt"
)

func main() {
	var card string = "Ace of Spades"
	fmt.Println(card)
	var card2 = "Ace of Spades"
	fmt.Println(card2)
	card3 := "Ace of Spades"
	fmt.Println(card3)
	card3 = "Joker"
	fmt.Println(card3)
	pi := 3.14
	fmt.Printf("%T\n", pi)

	card4 := newCard()
	fmt.Println(card4)

	cards := []string{"Five of Diamonds", "Ace of Spades", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	for i, v := range cards {
		fmt.Println(i, v)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
