package main

import "fmt"

type deck [] string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Chase"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			append(cards, {cardSuits, cardValues})
		}
	}
}

func (d deck) print() {
	for i, card:= range d {
		fmt.Println(i, card)
	}
}

func deal(d deck,handSize int) (deck, deck) {
	retun d[:handSize], d[handSize:]
}

func (d deck) toString() string {

}