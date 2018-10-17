package main

import "fmt"

func main(){
	//cards := cards.newDeck()
	//cards.print()

	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}

// Need to define our expected return value (Explicit typing)
//func newCard() string {
//	return "Five of Diamonds"
//}

// Variables can be declared outside of func(), however, they can't be assigned outside of func

// Array - Fixed length list
// Slice - An array of dynamic length

// Both arrays and slices must be of identical types


// understanding func WriteFile(filename string, data []byte, perm os.FileMode) error
//  data = raw data written to hard drive, slice of bytes (byte slice)
// perm = permissions used to create that file, if we need to create it

// a byte slice, think of a string, ascii character code

//Type conversion
