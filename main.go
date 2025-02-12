package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)
	cardDeck()
}

// The data type returned by the function should be mentioned
// after the function name (after the parantheses)
func newCard() string{
	return "Five of Diamonds"
}

func slices() {
	// Slice is a list that can grow or shrink
	cards := []string {"Queen of Hearts", newCard()}
	cards = append(cards, "Six of Clubs")
	for _, card := range cards {
		fmt.Println(card)
	} 
}

func cardDeck() {
	// cards := deck {"Queen of Hearts", newCard()}
	cards := newDeck()
	hand, cards := cards.deal(3)
	hand.print()
	cards.print()
}

func variableInitialization() {
	// This is the longer, boring way to initialize a variable
	var card string = "Ace of Spades"

	// This is the type inferenced way, but you can only do this inside a function, 
	// not at the package level
	yelp := "Yelp!"
	fmt.Println(card, yelp)

	// You can also do it like this, even at package level
	var delcareFirst int
	delcareFirst = 10
	fmt.Println(delcareFirst)
}