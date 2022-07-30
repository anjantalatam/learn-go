package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	anjan := person{
		firstName: "Anjan",
		lastName:  "Talatam",
		contact: contactInfo{
			email:   "anjan@talatam.com",
			zipCode: 123456,
		},
	}

	fmt.Printf("%+v", anjan)
}
