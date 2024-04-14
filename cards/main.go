package main

import "fmt"

func main() {

	myDeck := newDeck()
	fmt.Println("Full Deck")
	myDeck.print()

	fmt.Println("Shuffling")
	myDeck.shuffle()
	myDeck.print()

	hand, remainingCards := deal(myDeck, 5)

	fmt.Println("Dealt Hand")
	hand.print()

	fmt.Println("Remaining Deck")
	remainingCards.print()

	hand.saveToFile("my_cards.txt")
	persistedHand := newDeckFromFile("my_cards.txt")
	persistedHand.print()

}
