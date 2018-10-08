package main

import "fmt"

type deck[] string

/*
 Receiver function -
 d - similar to reference "this" / "self"
 actual copy of deck (convention of GO, first letter of "deck")
 deck - call this function on itself
*/

/**
 * Creates a new deck of two different card arrays.
 */
func newDeck() deck {
	cards:= deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//Replace unused variables with _
	for _, shape:= range cardSuits {
		for _, value:= range cardValues {
			cards = append(cards, shape + " of " + value)
		}
	}
	return cards
}

/**
 * Print receiver function on deck
 * Takes in a deck and prints iterative value over it
 */
func (d deck) print() {
	for i, card:= range d{
		fmt.Println(i, card)
	}
}

/**
 * Takes one deck and returns two different decks based on a given handsize
 */
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

