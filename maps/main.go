package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blud":  "#0000FF",
	}
	colors["white"] = "#FFFFFF"
	printMap(colors)
	delete(colors, "white")
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
