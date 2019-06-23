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

func (p person) updateFirstname(newName string) {
	p.firstname = newName;
}

func main() {
	contactInfo1 := contactInfo{email: "bob@gmail.com"}

	person1 := person{firstname: "bob", lastname: "Jones", contactInfo: contactInfo1}

	fmt.Println(person1);
	person1.display()

	person1.updateFirstname("jack")
	person1.display()
}
