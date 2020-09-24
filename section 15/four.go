package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %d years old \n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Natasha",
		last:  "Romanoff",
		age:   30,
	}
	p1.speak()
}
