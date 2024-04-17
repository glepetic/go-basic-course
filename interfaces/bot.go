package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
