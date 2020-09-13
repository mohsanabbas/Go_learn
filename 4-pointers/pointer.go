package main

import "fmt"

type contactInfo struct {
	email   string
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
			email:   "mohsanabbas@cvccorp.com.br",
			zipCode: 01214100,
		},
	}
	// jimPointer := &jim
	// another way to get memory address of jim  but if we don't want use to use this approach all we need to do is in reciever use  pointer
	jim.updateName("Ali")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

// reciver will recieve a person struct and turn it into pointer which is reference to original jim struct,
// In golang when we use reciever or functions it clones the data into different memory loation,
// so to access that we use pointers which get the memory address and point to it's value

func (p *person) updateName(newName string) {
	(*p).firstName = newName
	fmt.Printf("%+v", p)
}
