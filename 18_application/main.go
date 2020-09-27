package main

import (
	"encoding/json"
	"fmt"
)

// https://mholt.github.io/json-to-go/
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Captain","Last":"America","Age":72},{"First":"Natasha","Last":"Romanoff","Age":32}]`
	bs := []byte(s) //[]uint8

	var people []person // people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println("\n person at index", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
