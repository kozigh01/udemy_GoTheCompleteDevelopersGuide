package main

import (
	"fmt"

	"github.com/kozigh01/udemy_GoTheCompleteDevelopersGuide/structs/examples"
)

func main() {
	// struct01()

	struct02()

	struct03()
}

func struct01() {
	p1 := examples.Person {"Alex", "Anderson"}
	fmt.Println(p1)

	p2 := examples.Person {
		FirstName: "Bob",
		LastName: "Smith",
	}
	fmt.Println(p2)
	fmt.Printf("%q\n", p2)
	fmt.Printf("%+v\n", p2)

	var p3 examples.Person
	fmt.Printf("%q\n", p3)
	fmt.Printf("%+v\n", p3)

	p3.FirstName = "Martha"
	p3.LastName = "Wayne"
	fmt.Printf("%q\n", p3)
	fmt.Printf("%+v\n", p3)
}

func struct02() {
	bob := examples.Person2 {
		FirstName: "Bob",
		LastName: "Barker",
		Contact: examples.ContactInfo{
			Email: "a@b.com",
			Zip: 12345,
		},
	}

	fmt.Printf("%+v\n", bob)
}

func struct03() {
	billy := examples.Person3 {
		FirstName: "Billy",
		LastName: "Kidd",
		ContactInfo: examples.ContactInfo{
			Email: "d@e.com",
			Zip: 98765,
		},
	}
	billy.Print()

	billy.UpdateName("bobby")
	billy.Print()
	
	(&billy).UpdateName2("bobby")
	billy.Print()

	billy.UpdateName2("bubby")
	billy.Print()

	(&billy).UpdateName("carl")
	billy.Print()
}