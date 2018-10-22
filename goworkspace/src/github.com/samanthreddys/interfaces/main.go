package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	return "Hello!"

}

func (sb spanishBot) getGreeting() string {
	return "Hola!!"

}
func printGreeting(b bot) {
	fmt.Printf("%T", b)
	fmt.Println(b.getGreeting())
}
