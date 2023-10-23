package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hi, There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Interfaces
// Interfaces are 'implicit'
// Interfaces are a contract to help us manage types

// Maps
// All keys must be the same type
// All values must be the same type
// Keys are indexed - we can iterate over them
// Represents a collection of related properties

// Structs
// Values can be of different type
// Keys don't support indexing
// You need to know all the different fields at compile time
// Represents a "thing" with a lot of different properties
