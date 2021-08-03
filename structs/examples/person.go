package examples

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

type Person2 struct {
	FirstName string
	LastName string
	Contact ContactInfo
}

type Person3 struct {
	FirstName string
	LastName string
	ContactInfo
}

func (p Person3) Print() {
	fmt.Printf("%+v\n", p)
}