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

func createPerson(first string, last string, email string, zip int) person {
	return person{
		firstName: first,
		lastName:  last,
		contactInfo: contactInfo{
			email:   email,
			zipCode: zip,
		},
	}
}

func (p person) print() {
	fmt.Printf("%#v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	alex := createPerson("Alex", "Anderson", "hello@yahoo.com", 11024)
	alex.updateName("Alexia")
	alex.print()
}
