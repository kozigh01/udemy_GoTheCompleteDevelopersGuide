package langbot2

import "fmt"

type Bot interface {
	GetGreeting() string
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) GetGreeting() string {
	// custom english logic goes here
	return "Hi There"
}

func (SpanishBot) GetGreeting() string {
	// custom spanish logic goes here
	return "Hola"
}
