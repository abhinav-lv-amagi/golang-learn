package main

import "fmt"

/*
 * Create a new type of 'deck'
 * which is a slice of strings
 */
type deck []string

/*
 * Receiver function:
 * func (d deck) print()
 * 'd' is the actual copy of the deck we're working with
 * 'deck' => so that every variable of type 'deck' can call this function on itself
 * Eg: cards.print() (where 'cards' is of type deck)
 *
 * Note: The receiver argument is abbreviated as a single letter as a convention in Go
 */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Generates a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit + " of " + cardValue)
		}
	}

	return cards
}

/*
 * Function to deal a hand, so basically split the deck into two
 * There are two functions below for the same purpose, the first of 
 * which is my implementation. The second one uses:
 * 
 * Multiple Return Values - yes, Go functions can return multiple values.
*/
func (d deck) dealMyImpl(handSize int) []deck {
	return []deck{d[:handSize], d[handSize:]}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
