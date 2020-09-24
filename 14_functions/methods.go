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

func (a avenger) speak() {
	fmt.Println("I am", a.firstName, a.lastName)
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
	firstAvenger.speak()
	a2.speak()
	fmt.Println(firstAvenger)
}
