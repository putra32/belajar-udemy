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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 98000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()

}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
