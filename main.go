package main

import (
	"fmt"
	"learning-go/deck-of-cards"
	httpproject "learning-go/http-project"
	interfaceassmt "learning-go/interface-assmt"
	"learning-go/interfaces"
	"learning-go/maps"
	"learning-go/structs"
)

func CardDeckProject() {
	// deck of cards project
	deck.SaveCardDeck()
	deck.TestCardDeck()
}

func main() {
	learn := 6
	switch learn {
	case 1:
		CardDeckProject()
	case 2:
		structs.StructEntry()
	case 3:
		maps.MapEntry()
	case 4:
		interfaces.InterfaceEntry()
	case 5:
		httpproject.HttpEntry()
	case 6:
		interfaceassmt.InterfaceAssmtEntry()
	default:
		fmt.Println("Invalid option")
	}
}
