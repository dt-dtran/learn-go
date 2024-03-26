package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// both lines below is the same
	// contact   contactInfo
	contactInfo
}

func main() {
	// 1. struct literal: generic
	// diana := person{"Diana", "Tran"}

	// 2. struct literal: field Name: syntax
	// diana := person{firstName: "Diana", lastName: "Tran"}

	// 3. struct literal: zero value syntax with type struct
	// var diana person
	// diana.lastName = "T"
	// fmt.Println(diana.lastName)
	// fmt.Printf("%+v", diana) // provide entire structure of struct

	// 4. embed struct:
	jim := person{
		firstName: "Jim",
		lastName:  "Johns",
		// if field = contact contactInfo
		// contact: contactInfo{
		// 	email:   "jim@email.com",
		// 	zipCode: 001122,
		// },

		// if field = contactInfo
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 001122,
		},
	}
	// fmt.Println(jim)
	// fmt.Printf("%+v", jim)

	// 5. receiver
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
