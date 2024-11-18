package main

import "fmt"

var (
	isEmptyMap map[int]int
)

func main() {
	// empty maps
	var ok map[string]bool
	intOk := make(map[string]int)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)

	printEmptyMap(ok)
	printEmptyMap(intOk)
	printEmptyMap(isEmptyMap)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// generics scenario
func printEmptyMap[k comparable, v bool | int](m map[k]v) {
	fmt.Println(m)
}
