package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	person1 := person{"Anjan", "Talatam"}
	persons := []string{"Anjan", "talatam"}
	person2 := person{firstName: "Anjan", lastName: "Talatam"}
	fmt.Println(person1, person2, persons)
}
