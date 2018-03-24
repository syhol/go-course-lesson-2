package main

import (
	"fmt"
)

// contactinfo represents information for a person
type contactinfo struct {
	email   string
	zipCode int
}

// Person represents a real person
type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	constructionExample()
}

func constructionExample() {
	var p1 person
	p1.print()

	p2 := person{"Bob", "Smith", contactinfo{}}
	p2.print()

	p3 := person{
		firstName: "Bob",
		lastName:  "Smith",
		contact: contactinfo{
			email:   "bob@example.com",
			zipCode: 123456,
		},
	}
	p3.print()

	var p4 person
	p4.firstName = "Bob"
	p4.lastName = "Smith"
	p4.print()

	p4.updateName("Jim")
	p4.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
