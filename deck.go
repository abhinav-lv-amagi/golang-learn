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
