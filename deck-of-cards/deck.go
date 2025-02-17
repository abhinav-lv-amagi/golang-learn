package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* VARIABLES AND INITIALIZATION #################################
 */
func VariableInitialization() {
	// This is the longer, boring way to initialize a variable
	var card string = "Ace of Spades"

	// This is the type inferenced way, but you can only do this inside a function, 
	// not at the package level
	yelp := "Yelp!"
	fmt.Println(card, yelp)

	// You can also do it like this, even at package level
	var delcareFirst int = 10
	fmt.Println(delcareFirst)
}

/* SLICES #######################################################
 * Create a new type of 'deck'
 * which is a slice of strings
 */
type deck []string

func Slices() {
	// Slice is a list that can grow or shrink
	cards := []string {"Queen of Hearts", NewCard()}
	cards = append(cards, "Six of Clubs")
	for _, card := range cards {
		fmt.Println(card)
	} 
}

/* FUNCTIONS, MULTIPLE RETURN VALUES ############################
 */
// The data type returned by the function should be mentioned
// after the function name (after the parantheses)
func NewCard() string{
	return "Five of Diamonds"
}

/*
 * Function to deal a hand, so basically split the deck into two
 * There are two functions below for the same purpose, the first of 
 * which is my implementation. The second one uses:
 * 
 * Multiple Return Values - yes, Go functions can return multiple values.
*/
func (d deck) DealMyImpl(handSize int) []deck {
	return []deck{d[:handSize], d[handSize:]}
}


/* RECEIVER FUNCTIONS ###########################################
 * func (d deck) print()
 * 'd' is the actual copy of the deck we're working with
 * 'deck' => so that every variable of type 'deck' can call this function on itself
 * Eg: cards.print() (where 'cards' is of type deck)
 *
 * Note: The receiver argument is abbreviated as a single letter as a convention in Go
 */
func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/* FILE OPERATIONS ##############################################
 */
func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		  fmt.Println("Error: ", err)
		  os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

/* GENERATE RANDOM VALUES #######################################
 */
 func (d deck) Shuffle() {
	for i := range d {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// swapIdx := r.Int() % len(d) // Instead we can do this:
		swapIdx := r.Intn(len(d))
		d[i], d[swapIdx] = d[swapIdx], d[i]
	}
}

/* MISC FUNCTIONS FROM CARD PROJECT
 */
// Generates a new deck of cards
func NewDeck() deck {
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

func Deal(handSize int, d deck) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func SaveCardDeck() {
	// get new deck and save to file
	cards := NewDeck()
	cards.SaveToFile("deck-of-cards/bytecards.txt")
}

func TestCardDeck() {
	cards := NewDeckFromFile("deck-of-cards/bytecards.txt")
	cards.Print()
	cards.Shuffle()
	fmt.Println("------------")
	cards.Print()
}

