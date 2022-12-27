package main

import "fmt"

func main() {
	colors := map[string]interface{}{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(s map[string]interface{}) {
	for color, hex := range s {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}
