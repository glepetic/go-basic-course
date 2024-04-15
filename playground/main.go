package main

import "fmt"

/*
	Everything in Go is pass by value,
	but some types work like pass by reference due to how said types require auxiliar structures

	Types passed by value (consider when to use pointers)
	- int
	- float
	- string
	- bool
	- struct

	Types passed by reference (don't worry about pointers)
	- slice
	- map
	- channel
	- pointer
	- function
*/

func main() {

	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateMySlice(mySlice)

	fmt.Println(mySlice)

}

func updateMySlice(s []string) {
	/* The slice is copied by value, but only the slice header, not the array which the slice header points to
	   The latter remains the same, so this function effectively updates the value
	*/
	s[0] = "Bye"
}
