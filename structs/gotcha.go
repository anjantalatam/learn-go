package main

import "fmt"

func sliceUpdate() {
	mySlice := []string{"Hi", "There", "How", "r", "u"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
