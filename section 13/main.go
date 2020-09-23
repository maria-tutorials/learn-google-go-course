package main

import "fmt"

func main() {
	one()
}

func one() {
	type person struct {
		firstName   string
		lastName    string
		favIceCream []string
	}

	p1 := person{
		firstName: "Captain",
		lastName:  "America",
		favIceCream: []string{
			"Banana",
			"Chocolate",
		},
	}

	p2 := person{
		firstName: "Natasha",
		lastName:  "Romanof",
		favIceCream: []string{
			"Strawberry",
			"Raspberry",
			"Baunilha",
		},
	}

	fmt.Println("First person: \n", p1)
	fmt.Println(p1.firstName, p1.lastName)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println("Second person: \n", p1)
	fmt.Println(p2)
	fmt.Println(p2.firstName, p2.lastName)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}
