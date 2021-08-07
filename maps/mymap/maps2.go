package mymap

import (
	"fmt"
)

func Maps2() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	colors["10"] = "#101010"
	fmt.Printf("colors - type: %T, value: %v\n", colors, colors)

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Printf("Color: %v is hex code %v\n", color, hexCode)
	}
}