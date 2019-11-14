package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func printGreeting(eb englishbot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishbot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishbot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishbot) getGreeting() string {
	return "Hola!"
}
