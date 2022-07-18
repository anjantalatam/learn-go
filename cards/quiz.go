package main

import "fmt"

func quiz() {
	color1, color2, color3 := colors()

	fmt.Println(color1, color2, color3)
}

func colors() (string, string, string) {
	return "red", "yellow", "blue"
}
