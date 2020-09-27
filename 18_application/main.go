package main

import (
	"encoding/json"
	"fmt"
)

//NOTE: lowercase
type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Captain",
		last:  "America",
		age:   72,
	}
	p2 := person{
		first: "Natasha",
		last:  "Romanoff",
		age:   32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
