package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type kambaBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	kb := kambaBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(kb)
}

func (kambaBot) getGreeting() string {
	return "Uvoo"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Holla"
}
