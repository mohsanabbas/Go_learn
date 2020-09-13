package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo // can replace with just type contactInfo
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
	fmt.Printf("%+v", jim)

}
