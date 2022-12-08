package main

import "fmt"

type bot interface {
	greetingBot() string
}

type englishBot struct{}
type spainBot struct{}

func main() {
	eb := englishBot{}
	sb := spainBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.greetingBot())
}

func (eb englishBot) greetingBot() string {
	return "Hi, There!"
}

func (sp spainBot) greetingBot() string {
	return "Holla!"
}
