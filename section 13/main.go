package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, ":")
		fmt.Println(v.firstName, v.lastName)

		for i, v := range v.favIceCream {
			fmt.Println(i, v)
		}
	}
}
