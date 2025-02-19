package interfaces

import "fmt"

type englishBot struct{}
type spanishBot struct{}

/*
 * To whom it may concern, our program has a new type called 'bot'
 * If you are a type in this program with a function called `getGreeting`,
 * and it returns a string then you are now an honorary member of type `bot`.
 * Now, you can call any function that accepts a type `bot` with your type.
 */
type bot interface {
	getGreeting() string

	// If we wanted to mention arg types and multiple return types:
	// getGreeting(string, int) (string, error)
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func InterfaceEntry() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
