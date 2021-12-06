package main

import (
	"fmt"
)

func main() {
	//a map is a reference type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	delete(colors, "red")
	changeMap(colors, "black", "#000000")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hexcode := range c {
		fmt.Println("Hex code for", color, "is", hexcode)
	}
}

func changeMap(c map[string]string, color string, hexcode string) {
	c[color] = hexcode
}
