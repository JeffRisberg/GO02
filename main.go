package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

func main() {
	person1 := person{"bob", "Jones"}
	fmt.Println(person1);

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
