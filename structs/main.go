package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	// person1 := person{"Anjan", "Talatam"} //1
	// person2 := person{firstName: "Anjan", lastName: "Talatam"} //2

	var person3 person //3

	person3.firstName = "Anjan"
	person3.lastName = "Talatam"

	fmt.Println(person3)
	fmt.Printf("%+v", person3)

}
