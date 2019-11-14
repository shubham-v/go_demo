package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0001",
		"blue":  "#ff0002",
	}

	colors["white"] = "#fffff"

	// 	delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}
