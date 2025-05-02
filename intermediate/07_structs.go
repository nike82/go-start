package intermediate

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Adderss
	PoneHomeCell
}

type Adderss struct {
	city   string
	contry string
}

type PoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	p1 := Person{
		firstName: "John",
		lastName:  "Full",
		age:       30,
		address: Adderss{
			city:   "bohn",
			contry: "Ge",
		},
		PoneHomeCell: PoneHomeCell{
			home: "1242",
			cell: "4376484586",
		},
	}

	p2 := Person{
		firstName: "Bob",
		age:       55,
	}

	p2.address.city = "NY"
	p2.address.contry = "NY"

	p3:=Person{
		firstName: "Bob",
		age:       55,
	}

	p3.address.city = "NY"
	p3.address.contry = "NY"

	fmt.Println(p1.firstName)
	fmt.Println(p1.address.city)
	fmt.Println(p1.address.contry)
	fmt.Println(p1.cell)
	fmt.Println(p2.firstName)
	fmt.Println(p2 == p1)
	fmt.Println(p2 == p3)



	user := struct {
		userName string
		email    string
	}{
		userName: "Jane",
		email:    "jane11@mail.com",
	}

	fmt.Println(user.userName)
	fmt.Println(p1.fullName())

	fmt.Println(p1.age)
	p1.incrementAgeByOne()
	fmt.Println(p1.age)
}
