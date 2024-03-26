package main

import "fmt"

// 2. with interface
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello! - English Bot"
}

func (sb spanishBot) getGreeting() string {
	return "Hola! - Bot español"
}

// 1. without interface
// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

	// printGreeting(eb)
	// without interface a separate function is needed for each greeting.
	// printGreetingSB(sb)
// }

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreetingSB(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// func (eb englishBot) getGreeting() string {
// 	return "Hello! - English Bot"
// }

// func (sb spanishBot) getGreeting() string {
// 	return "Hola! - Bot español"
// }
