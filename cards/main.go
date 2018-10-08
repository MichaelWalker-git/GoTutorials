package main

import "fmt"

func main(){
	//var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
}

// Need to define our expected return value (Explicit typing)
func newCard() string {
	return "Five of Diamonds"
}

// Variables can be declared outside of func(), however, they can't be assigned outside of func

// Array - Fixed length list
// Slice - An array of dynamic length

// Both arrays and slices must be of identical types

