package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) introduce() {
	fmt.Println(p.firstName, p.lastName, "is", p.age, "years old")
}

func (p person) print() {
	fmt.Printf("%+v\n", p) // %+v prints the field names with their corresponding values
}

func (p person) updateNameByValue(newFirstName string) person {
	p.firstName = newFirstName
	return p
}

func (pointerToP *person) updateNameByReference(newFirstName string) {
	// In Go course, instructor used (*pointerToP).firstName = newFirstName to turn the address/pointer into a value
	pointerToP.firstName = newFirstName
}
