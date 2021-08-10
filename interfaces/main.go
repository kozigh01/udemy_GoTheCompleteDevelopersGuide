package main

import (
	"fmt"

	// "github.com/mkozigo/kozigh01/udemy_GoTheCompleteDevelopersGuide/interfaces/interfaces"
	"github.com/mkozigo/kozigh01/udemy_GoTheCompleteDevelopersGuide/interfaces/langbot1"
	"github.com/mkozigo/kozigh01/udemy_GoTheCompleteDevelopersGuide/interfaces/langbot2"
)

func main() {
	fmt.Println("hello from interfaces")

	fmt.Println("Language Bot 1")
	eb := langbot1.EnglishBot{}
	fmt.Println(eb.GetGreeting())
	langbot1.PrintGreeting(eb)

	sb := langbot1.SpanishBot{}
	fmt.Println(sb.GetGreeting())
	sb.PrintGreeting()

	fmt.Println("Language Bot 2")
	eb2 := langbot2.EnglishBot{}
	sb2 := langbot2.SpanishBot{}

	langbot2.PrintGreeting(eb2)
	langbot2.PrintGreeting(sb2)

	var test1 langbot2.Bot = langbot2.EnglishBot{}
	langbot2.PrintGreeting(test1)
}