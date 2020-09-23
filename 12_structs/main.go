package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   62,
	}

	fmt.Println(p1)

	fmt.Println(p1.first, p1.last, p1.age)
}
