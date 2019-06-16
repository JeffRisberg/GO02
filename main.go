package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func (p person) display() {
	fmt.Println(p.firstname + p.lastname)
	fmt.Println(p.contactInfo.email)
}

func main() {
	contactInfo1 := contactInfo{email: "bob@gmail.com"}

	person1 := person{firstname: "bob", lastname: "Jones", contactInfo: contactInfo1}

	fmt.Println(person1);
	person1.display()
	/*
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("")
	remainingDeck.print()

	fmt.Println(hand.toString())

	hand.saveToFile("hand20190609")

	hand2 := newDeckFromFile("hand20190609")

	hand2.print()
	*/
}
