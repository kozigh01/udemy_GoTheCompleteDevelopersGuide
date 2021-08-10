package langbot1

import "fmt"

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) GetGreeting() string {
	// custom english logic goes here
	return "Hi There"
}

func PrintGreeting(eb EnglishBot) {
	fmt.Println(eb.GetGreeting())
}

func (_ SpanishBot) GetGreeting() string {
	// custom spanish logic goes here
	return "Hola"
}

func (sb SpanishBot) PrintGreeting() {
	fmt.Println(sb.GetGreeting())
} 