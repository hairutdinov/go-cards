package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Туз Пик
	// card := newCard()
	// fmt.Println(card)
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining deck:")
	remainingDeck.print()
	// cards.print()
}
