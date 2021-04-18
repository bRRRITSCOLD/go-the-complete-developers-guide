package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	p1 := person{"Alex", "Anderson", contact{"test@email.com", 55555}}
	p1.updateFirstName("Jim")
	p1.print()

	p2 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contact{
			email: "test@email.com",
			zip:   55555,
		},
	}
	p2.print()

	var p3 person
	p3.firstName = "Alex"
	p3.lastName = "Anderson"
	p3.contact.email = "test@email.com"
	p3.contact.zip = 55555
	p3.print()
}
