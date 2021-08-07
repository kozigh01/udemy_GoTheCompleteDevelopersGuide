package mymap

import "fmt"

func Maps1() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}
	colors["white"] = "#ffffff"
	fmt.Printf("colors - type: %T, value: %v\n", colors, colors)
	fmt.Printf("colors[\"blue\"] - type: %T, value: %v\n", colors["blue"], colors["blue"])
	delete(colors, "green")
	fmt.Printf("colors - type: %T, value: %v\n", colors, colors)
	
	var colors2 map[string]int
	fmt.Printf("colors - type: %T, value: %v\n", colors2, colors2)
	
	var colors3 = make(map[string]float64)
	fmt.Printf("colors - type: %T, value: %v\n", colors3, colors3)
}