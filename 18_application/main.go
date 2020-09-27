package main

import (
	"encoding/json"
	"fmt"
)

//NOTE: uppercase
type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Captain",
		Last:  "America",
		Age:   72,
	}
	p2 := person{
		First: "Natasha",
		Last:  "Romanoff",
		Age:   32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
