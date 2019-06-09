package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("")
	remainingDeck.print()

	fmt.Println(hand.toString())

	hand.saveToFile("hand20190609")

	hand2 := newDeckFromFile("hand20190609")

	hand2.print()
}
