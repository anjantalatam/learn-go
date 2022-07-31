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
	// updateLastName(anjan, "t") // throws error cannot use anjan (variable of type person) as *person value in argument
	updateLastName((anjanPointer), "SP")

	anjan.print()
	sliceUpdate()

}

func (pointerToPerson *person) updateName(newUpdatePerson string) {
	(*pointerToPerson).firstName = newUpdatePerson
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateLastName(pp *person, newLastName string) {
	(*pp).lastName = newLastName

}
