package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func (b englishBot) getGreeting() string {
	return "Hello World"
}

func (b spanishBot) getGreeting() string {
	return "Hola mundo"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	printGreeting(englishBot{})
	printGreeting(spanishBot{})
}
