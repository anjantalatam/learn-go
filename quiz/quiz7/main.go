// What will the following program print out?

package main

import "fmt"

func main() {
	name := "Bill"

	fmt.Println(*&name)
}
