package main

import (
	"fmt"
)

// Person represents a Person
type Person struct {
	firstname string
	lastname  string
}

// Person2 represents a Person
type Person2 struct {
	firstname string
	lastname  string
	contact   ContactInfo
}

// Print prints the Person2
func (p Person2) Print() {
	fmt.Printf("Hello, my name is %v %v and my email is %v\n", p.firstname, p.lastname, p.contact.email)
}

func (p *Person2) updateName(n string) {
	p.lastname = n
}

// ContactInfo encapsulates contact information
type ContactInfo struct {
	email string
	zip   int
}

func main() {
	fmt.Println(Person{"Christian", "Dietrich"})
	fmt.Println(Person{firstname: "Christian", lastname: "Dietrich"})
	fmt.Println(Person{
		firstname: "Christian",
		lastname:  "Dietrich",
	})
	// anonymous
	fmt.Println(struct {
		firstname string
		lastname  string
	}{"Christian", "Dietrich"})

	var me Person
	fmt.Println(me) // zero values
	fmt.Printf("%q\n", me)
	me.firstname = "Christian"
	me.lastname = "Dietrich"
	fmt.Println(me)
	fmt.Printf("%q %+v\n", me, me)

	me2 := Person2{
		firstname: "Christian",
		lastname:  "Dietrich",
		contact: ContactInfo{
			email: "sample@sample.com",
			zip:   12345,
		},
	}
	fmt.Println(me2)
	fmt.Printf("%#v\n", me2)

	me2.Print()
	me2.updateName("Bond")
	me2.Print()
	me2ptr := &me2
	me2ptr.updateName("Christian")
	me2.Print()
}
