package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo //  just Type contactInfo or contact contactInfo
}

func main() {

	jim := person{
		firstName: "Mohsan",
		lastName:  "Abbas",
		contactInfo: contactInfo{
			address: "avenida duque de caxia 61",
			zipCode: 01214100,
		},
	}
	jim.print()
	jim.updateName("Ali")
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

// this will lead us to learn about go pointers, because when we are updating name which is not same memory location it will not effect the update.
func (p person) updateName(newName string) {
	p.firstName = newName

}
