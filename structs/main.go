package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	anjan := person{
		firstName: "Anjan",
		lastName:  "Talatam",
		contactInfo: contactInfo{
			email:   "anjan@talatam.com",
			zipCode: 123456,
		},
	}

	anjanPointer := &anjan
	fmt.Printf("%p", anjanPointer)
	anjanPointer.updateName("Surya")
	anjan.print()

}

func (pointerToPerson *person) updateName(newUpdatePerson string) {
	(*pointerToPerson).firstName = newUpdatePerson
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
