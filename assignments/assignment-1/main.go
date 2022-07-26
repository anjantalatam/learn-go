package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		var numType string
		if num%2 == 0 {
			numType = "even"
		} else {
			numType = "odd"
		}
		// Both the below lines are same
		fmt.Printf("%d is %s\n", num, numType)
		fmt.Println(num, "is", numType)
	}
}
