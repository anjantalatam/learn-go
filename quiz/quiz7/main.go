// What does the output of the code?
// Bill
// Alex
// Empty String
package main

import "fmt"

func main() {
	name := "Bill"
	updateValue(name)
	fmt.Println(name)
}

func updateValue(n string) {
	n = "Alex"
}
