package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Not Maria"
}

func main() {
	p := person{
		name: "Maria",
	}
	fmt.Println(p)

	changeMe(&p)
	fmt.Println(p)
}
