package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   62,
		},
		ltk: true,
	}

	p2 := person{
		first: "Also",
		last:  "Bond",
		age:   50,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
