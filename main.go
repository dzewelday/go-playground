package main

import "fmt"

func main() {

	person := createPerson()
	person.print()
}

type Person struct {
	firstName  string
	middleName string
	lastName   string
}

func createPerson() Person {
	return Person{
		firstName:  "Dawit",
		middleName: "G",
		lastName:   "Zewelday",
	}
}

func (p Person) print() {
	fmt.Printf("Hey %s %s %s", p.firstName, p.middleName, p.lastName)
}
