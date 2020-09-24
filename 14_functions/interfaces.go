package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type avenger struct {
	person
	powers bool
}

type human interface {
	speak()
}

func (a avenger) speak() {
	fmt.Println("I am", a.firstName, a.lastName, " - person spoke")
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "- AVENGER spoke")
}

func drink(h human) {
	fmt.Println("Im inside bar", h)
}

func main() {
	firstAvenger := avenger{
		person: person{
			"Captain",
			"America",
		},
		powers: true,
	}
	a2 := avenger{
		person: person{
			"Natasha",
			"Romanoff",
		},
		powers: true,
	}

	p1 := person{
		firstName: "agent",
		lastName:  "Phil",
	}

	firstAvenger.speak()
	a2.speak()
	p1.speak()

	drink(firstAvenger)
	drink(a2)
	drink(p1)
}
