package main

import "fmt"

func main() {
	// var colors map[string]string // 2

	colors := make(map[int]string) // 3

	// colors := map[string]string{ // 1
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	// Insert a key,value pair
	colors[10] = "#ffffff"

	// Delete a key,value pair
	delete(colors, 10)

	// 11 doesn't exist in colors: So colors is unaltered with this below line
	delete(colors, 11)

	fmt.Println(colors)
}
