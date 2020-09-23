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

	s := struct {
		name   string
		powers map[string]int
	}{
		name: p1.firstName + " " + p1.lastName,
		powers: map[string]int{
			"Shield":  10,
			"Muscles": 9,
		},
	}
	fmt.Println(s.name)
	for k, v := range s.powers {
		fmt.Println(k, v)
	}
}
