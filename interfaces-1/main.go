package main

import "fmt"

type Bot interface {
	getGreetings() string
}

func printGreetings(b Bot) {
	fmt.Println(b.getGreetings())
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreetings() string {
	return "hello"
}
func (spanishBot) getGreetings() string {
	return "hola"
}
func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreetings(eb)
	printGreetings(sb)
}
