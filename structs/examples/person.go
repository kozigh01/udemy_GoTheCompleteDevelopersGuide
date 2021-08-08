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

func (p Person3) UpdateName(newName string) {
	p.FirstName = newName
}

func (p *Person3) UpdateName2(newName string) {
	p.FirstName = newName
}